package http

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/elastos/Elastos.ELA.Utility/common"
	"github.com/elastos/Elastos.ORG.API.Misc/chain"
	"github.com/elastos/Elastos.ORG.API.Misc/config"
	"github.com/elastos/Elastos.ORG.API.Misc/db"
	"github.com/elastos/Elastos.ORG.API.Misc/log"
	"github.com/elastos/Elastos.ORG.API.Misc/tools"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

var (
	dba = db.NewInstance()
)

func StartServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	log.Infof("start server at: %s", config.Conf.ServerPort)
	err := http.ListenAndServe(":"+config.Conf.ServerPort, router)
	if err != nil {
		log.Fatal("Error start server :" + err.Error())
	}
}

//searchKey search did property key
func searchKey(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	did := params["did"]
	key := params["key"]

	c, err := dba.ToInt("select count(*) from chain_did_property where (did_status = 0 or property_key_status = 0) and did ='" + did + "' and property_key = '" + key + "'")
	if err != nil {
		w.Write([]byte(`{"result":"` + err.Error() + `","status":500}`))
		return
	}
	if c > 0 {
		w.Write([]byte(`{"result":"did is discarded or property key is discarded","status":200}`))
		return
	}
	v, err := dba.ToStruct("select Did,Did_status,Public_key,Property_key,property_value,txid,block_time,height from chain_did_property where did ='"+did+"' and property_key = '"+key+"' and did_status = 1 and property_key_status = 1 order by id desc limit 1", chain.Did_Property{})
	if err != nil {
		w.Write([]byte(`{"result":"` + err.Error() + `","status":500}`))
		return
	}
	if len(v) == 0 {
		w.Write([]byte(`{"status":200}`))
		return
	}
	b, err := json.Marshal(v[0])
	if err != nil {
		w.Write([]byte(`{"result":"` + err.Error() + `","status":500}`))
		return
	}
	w.Write([]byte(`{"result":` + string(b) + `,"status":200}`))
}

//history the address transaction history
func history(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	address := params["addr"]
	pageNum := r.FormValue("pageNum")
	f := r.FormValue("type")
	count := r.FormValue("showCount")
	showCount := false
	var err error
	if count != "" {
		showCount, err = strconv.ParseBool(count)
		if err != nil {
			w.Write([]byte(`{"result":"Invalid param , count should be a bool value","status":400}`))
			return
		}
	}
	order := r.FormValue("order")
	if order != "" && order != "desc" && order != "asc" {
		w.Write([]byte(`{"result":"Invalid param , order should only be desc or asc","status":400}`))
		return
	}
	if order == "" {
		order = "asc"
	}
	totalNum := 0
	bhs := make([]chain.Block_transaction_history, 0)
	if !showCount {
		var sql string
		if pageNum != "" {
			pageSize := r.FormValue("pageSize")
			var size int64
			if pageSize != "" {
				var err error
				size, err = strconv.ParseInt(pageSize, 10, 64)
				if err != nil {
					w.Write([]byte(`{"result":"` + err.Error() + `","status":400}`))
					return
				}
			} else {
				size = 10
			}
			num, err := strconv.ParseInt(pageNum, 10, 64)
			if err != nil {
				w.Write([]byte(`{"result":"` + err.Error() + `","status":400}`))
				return
			}
			if num <= 0 {
				num = 1
			}
			from := (num - 1) * size
			sql = "select address,txid,type,value,createTime,height,inputs,outputs,fee,txType,memo from chain_block_transaction_history where address = '" + address + "' order by id " + order + " limit " + strconv.FormatInt(from, 10) + "," + strconv.FormatInt(size, 10)
		} else {
			sql = "select address,txid,type,value,createTime,height,inputs,outputs,fee,txType,memo from chain_block_transaction_history where address = '" + address + "' order by id " + order
		}
		l, err := dba.Query(sql)
		if err != nil {
			w.Write([]byte(`{"result":"` + err.Error() + `","status":500}`))
			return
		}
		for e := l.Front(); e != nil; e = e.Next() {
			history := new(chain.Block_transaction_history)
			line := e.Value.(map[string]interface{})
			tools.Map2Struct(line, history)
			inputsArr := strings.Split(line["inputs"].(string), ",")
			outputsArr := strings.Split(line["outputs"].(string), ",")
			if f == "full" {
				history.Inputs = inputsArr[:len(inputsArr)-1]
				history.Outputs = outputsArr[:len(outputsArr)-1]
			} else {
				if history.Type == "income" {
					if len(inputsArr) > 0 {
						history.Inputs = []string{inputsArr[0]}
					} else {
						history.Inputs = []string{}
					}
					history.Outputs = []string{history.Address}
				} else {
					history.Inputs = []string{history.Address}
					history.Outputs = []string{history.Outputs[0]}
				}
			}
			if err != nil {
				w.Write([]byte(`{"result":"` + err.Error() + `","status":500}`))
				return
			}
			rawMemo, err := hex.DecodeString(history.Memo)
			if err != nil {
				w.Write([]byte(`{"result":"` + err.Error() + `","status":500}`))
				return
			}
			history.Memo = string(rawMemo)
			bhs = append(bhs, *history)
		}
	}

	l, err := dba.Query("select count(*) as count from chain_block_transaction_history where address = '" + address + "'")
	if err != nil {
		w.Write([]byte(`{"result":"` + err.Error() + `","status":500}`))
		return
	}
	totalNum, _ = strconv.Atoi(l.Front().Value.(map[string]interface{})["count"].(string))
	addrHis := chain.Address_history{bhs, totalNum}
	buf, err := json.Marshal(&addrHis)
	if err != nil {
		w.Write([]byte(`{"result":"` + err.Error() + `","status":500}`))
		return
	}
	w.Write([]byte(`{"result":` + string(buf) + `,"status":200}`))
}

func getPublicKey(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	addr := params["addr"]
	_, err := common.Uint168FromAddress(addr)
	if err != nil {
		http.Error(w, `{"result":"Invalid address","status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
		return
	}
	pub, err := dba.ToString("select publicKey from chain_block_transaction_history where address = '" + addr + "' and publicKey is not null and publicKey != '' limit 1")

	if pub == "" {
		w.Write([]byte(`{"result":"Can not find pubkey of this address, please using this address send a transaction first","status":200}`))
		return
	}
	w.Write([]byte(`{"result":"` + pub + `","status":200}`))
}

func producerStatistic(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pub := params["producer"]
	height := params["height"]
	if pub == "" || len(pub) != 66 {
		http.Error(w, `{"result":"invalid public key","status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
		return
	}
	type ret struct {
		Producer_public_key string `json:",omitempty"`
		Vote_type           string `json:",omitempty"`
		Txid                string `json:",omitempty"`
		Value               string `json:",omitempty"`
		Outputlock          int    `json:",omitempty"`
		Address             string `json:",omitempty"`
		Block_time          int64  `json:",omitempty"`
		Height              int64  `json:",omitempty"`
	}
	iHeight, err := strconv.Atoi(height)
	if err != nil {
		iHeight = 99999999
	}
	rst, err := dba.ToStruct("select Producer_public_key,Vote_type,Txid,Value,Address,Block_time,Height from chain_vote_info where producer_public_key = '"+pub+"' and (outputlock = 0 or outputlock >= height) and is_valid = 'YES' and height <= "+strconv.Itoa(iHeight), ret{})
	if err != nil {
		http.Error(w, `{"result":"internal error : `+err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
		return
	}
	buf, err := json.Marshal(&rst)
	if err != nil {
		http.Error(w, `{"result":"internal error : `+err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
		return
	}
	w.Write([]byte(`{"result":` + string(buf) + `,"status":200}`))
}

func voterStatistic(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	addr := params["address"]
	if addr == "" || len(addr) != 34 {
		http.Error(w, `{"result":"invalid address","status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
		return
	}
	pageNum := r.FormValue("pageNum")
	var sql string
	var from int64
	var size int64
	if pageNum != "" {
		pageSize := r.FormValue("pageSize")
		if pageSize != "" {
			var err error
			size, err = strconv.ParseInt(pageSize, 10, 64)
			if err != nil {
				w.Write([]byte(`{"result":"` + err.Error() + `","status":400}`))
				return
			}
		} else {
			size = 10
		}
		num, err := strconv.ParseInt(pageNum, 10, 64)
		if err != nil {
			w.Write([]byte(`{"result":"` + err.Error() + `","status":400}`))
			return
		}
		if num <= 0 {
			num = 1
		}
		from = (num - 1) * size
	}
	sql = "select * from chain_vote_info where address = '" + addr + "' order by id desc "
	info, err := dba.ToStruct(sql, chain.Vote_info{})
	if err != nil {
		http.Error(w, `{"result":"Internal error","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
		return
	}
	headersContainer := make(map[string]*chain.Vote_statistic_header)
	for i := 0; i < len(info); i++ {
		data := info[i].(*chain.Vote_info)
		h, ok := headersContainer[data.Txid+strconv.Itoa(data.N)]
		if ok {
			h.Node_num += 1
			h.Nodes = append(h.Nodes, data.Producer_public_key)
		} else {
			h = new(chain.Vote_statistic_header)
			h.Value = data.Value
			h.Node_num = 1
			h.Txid = data.Txid
			h.Height = data.Height
			h.Nodes = []string{data.Producer_public_key}
			h.Block_time = data.Block_time
			h.Is_valid = data.Is_valid
			headersContainer[data.Txid+strconv.Itoa(data.N)] = h
		}
	}
	var voteStatisticSorter chain.Vote_statisticSorter
	for _, v := range headersContainer {
		voteStatisticSorter = append(voteStatisticSorter, chain.Vote_statistic{
			*v,
			[]chain.Vote_info{},
		})
	}
	sort.Sort(voteStatisticSorter)
	if !(from == 0 && size == 0) && int(from+1+size) <= len(voteStatisticSorter) {
		voteStatisticSorter = voteStatisticSorter[from : from+size]
	} else if !(from == 0 && size == 0) && int(from+1) <= len(voteStatisticSorter) && int(from+1+size) > len(voteStatisticSorter) {
		voteStatisticSorter = voteStatisticSorter[from:]
	} else {
		voteStatisticSorter = chain.Vote_statisticSorter{}
	}
	var voteStatistic chain.Vote_statisticSorter
	ranklisthoder := make(map[int64][]interface{})
	//height+producer_public_key : index
	ranklisthoderByProducer := make(map[string]int)
	for _, _v := range voteStatisticSorter {
		v := _v.Vote_Header
		rst, ok := ranklisthoder[v.Height]
		if !ok {
			rst, err = dba.ToStruct(`select m.*,(@row_number:=@row_number + 1) as "rank" from (select ifnull(a.producer_public_key,b.ownerpublickey) as producer_public_key , ifnull(a.value,0) as value , b.* from 
(select A.producer_public_key , ROUND(sum(value),8) as value from chain_vote_info A where (A.cancel_height > `+strconv.Itoa(int(v.Height))+` or
 cancel_height is null) and height <= `+strconv.Itoa(int(v.Height))+` group by producer_public_key) a right join chain_producer_info b on a.producer_public_key = b.ownerpublickey 
 order by value desc) m ,  (SELECT @row_number:=0) AS t`, chain.Vote_info{})
			if err != nil {
				http.Error(w, `{"result":"internal error : `+err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
				return
			}
			totalVote, err := dba.ToFloat(`	select sum(a.value)  from (select A.producer_public_key , sum(value) as value from chain_vote_info A where (A.cancel_height > ` + strconv.Itoa(int(v.Height)) + ` or
	 cancel_height is null) and height <= ` + strconv.Itoa(int(v.Height)) + ` group by producer_public_key order by value desc limit 96) a`)
			if err != nil {
				http.Error(w, `{"result":"internal error : `+err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
				return
			}
			for _, r := range rst {
				vi := r.(*chain.Vote_info)
				addr, err := tools.GetAddress(vi.Ownerpublickey)
				if err != nil {
					log.Warn("Invalid Ownerpublickey " + vi.Ownerpublickey)
					continue
				}
				vi.Address = addr
				val, err := dba.ToString("select value from chain_block_transaction_history where height = " + strconv.Itoa(int(v.Height)) + " and txType = 'CoinBase' and value < " + strconv.Itoa(tools.Miner_Reward_PerBlock) + " and address = '" + addr + "'")
				if err != nil {
					log.Warn("Invalid Ownerpublickey " + vi.Ownerpublickey)
					continue
				}
				if val != "" {
					vi.Reward = val
				} else {
					vi.Reward = "0"
				}

				var vote float64
				if vi.Value == "" {
					vote = 0
				} else {
					vote, err = strconv.ParseFloat(vi.Value, 64)
					if err != nil {
						http.Error(w, `{"result":"internal error : `+err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
						return
					}
				}
				if vi.Rank <= 24 {
					vi.EstRewardPerYear = strconv.FormatFloat(float64(175834088*0.25/(100000000*36)*365*720+175834088*0.75/(totalVote*100000000)*vote*365*720), 'f', 8, 64)
				} else if vi.Rank <= 96 {
					vi.EstRewardPerYear = strconv.FormatFloat(float64(175834088*0.75/(totalVote*100000000)*vote*365*720), 'f', 8, 64)
				} else {
					vi.EstRewardPerYear = "0"
				}
			}
			for m := 0; m < len(rst); m++ {
				ranklisthoderByProducer[strconv.Itoa(int(v.Height))+rst[m].(*chain.Vote_info).Producer_public_key] = m
			}
		}
		var voteInfos []chain.Vote_info
		for _, pub := range v.Nodes {
			voteInfos = append(voteInfos, *rst[ranklisthoderByProducer[strconv.Itoa(int(v.Height))+pub]].(*chain.Vote_info))
		}
		voteStatistic = append(voteStatistic, chain.Vote_statistic{
			v,
			voteInfos,
		})
	}
	buf, err := json.Marshal(&voteStatistic)
	if err != nil {
		http.Error(w, `{"result":"internal error : `+err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
		return
	}
	w.Write([]byte(`{"result":` + string(buf) + `,"status":200}`))
}

func producerRankByHeight(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	height := params["height"]
	h, ok := strconv.Atoi(height)
	if ok != nil || h < 0 {
		http.Error(w, `{"result":"invalid height","status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
		return
	}
	state := r.FormValue("state")
	if state != "" && state != "active" && state != "inactive" && state != "pending" &&
		state != "canceled" && state != "illegal" && state != "returned" {
		http.Error(w, `{"result":"state can be one of the folowing values active,inactive,pending,canceled,illegal,returned","status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
		return
	}
	var err error
	var rst []interface{}
	if state == "" {
		rst, err = dba.ToStruct(`select m.*,(@row_number:=@row_number + 1) as "rank" from (select ifnull(a.producer_public_key,b.ownerpublickey) as producer_public_key , ifnull(a.value,0) as value , b.* from 
(select A.producer_public_key , ROUND(sum(value),8) as value from chain_vote_info A where (A.cancel_height > `+height+` or
 cancel_height is null) and height <= `+height+` group by producer_public_key) a right join chain_producer_info b on a.producer_public_key = b.ownerpublickey 
 order by value desc) m ,  (SELECT @row_number:=0) AS t `, chain.Vote_info{})
	} else {
		rst, err = dba.ToStruct(`select m.*,(@row_number:=@row_number + 1) as "rank" from (select ifnull(a.producer_public_key,b.ownerpublickey) as producer_public_key , ifnull(a.value,0) as value , b.* from 
(select A.producer_public_key , ROUND(sum(value),8) as value from chain_vote_info A where (A.cancel_height > `+height+` or
 cancel_height is null) and height <= `+height+` group by producer_public_key) a right join chain_producer_info b on a.producer_public_key = b.ownerpublickey where b.state = '`+strings.ToUpper(state[:1])+state[1:]+`'
 order by value desc) m ,  (SELECT @row_number:=0) AS t `, chain.Vote_info{})
	}
	if err != nil {
		http.Error(w, `{"result":"internal error : `+err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
		return
	}

	totalVote, err := dba.ToFloat(`	select sum(a.value)  from (select A.producer_public_key , sum(value) as value from chain_vote_info A where A.cancel_height > ` + height + ` or
	 cancel_height is null group by producer_public_key order by value desc limit 96) a`)
	if err != nil {
		http.Error(w, `{"result":"internal error : `+err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
		return
	}
	for _, r := range rst {
		vi := r.(*chain.Vote_info)
		addr, err := tools.GetAddress(vi.Ownerpublickey)
		if err != nil {
			log.Warn("Invalid Ownerpublickey " + vi.Ownerpublickey)
			continue
		}
		vi.Address = addr
		val, err := dba.ToString("select sum(value) from chain_block_transaction_history where txType = 'CoinBase' and address = '" + addr + "'")
		if err != nil {
			log.Warn("Invalid Ownerpublickey " + vi.Ownerpublickey)
			continue
		}
		if val != "" {
			iv, _ := strconv.Atoi(val)
			vi.Reward = strconv.FormatFloat(float64(iv)/100000000.0, 'f', 8, 64)
		} else {
			vi.Reward = "0"
		}

		var vote float64
		if vi.Value == "" {
			vote = 0
		} else {
			vote, err = strconv.ParseFloat(vi.Value, 64)
			if err != nil {
				http.Error(w, `{"result":"internal error : `+err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
				return
			}
		}
		if vi.Rank <= 24 {
			vi.EstRewardPerYear = strconv.FormatFloat(float64(175834088*0.25/(100000000*36)*365*720+175834088*0.75/(totalVote*100000000)*vote*365*720), 'f', 8, 64)
		} else if vi.Rank <= 96 {
			vi.EstRewardPerYear = strconv.FormatFloat(float64(175834088*0.75/(totalVote*100000000)*vote*365*720), 'f', 8, 64)
		} else {
			vi.EstRewardPerYear = "0"
		}
	}

	buf, err := json.Marshal(&rst)
	if err != nil {
		http.Error(w, `{"result":"internal error : `+err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
		return
	}
	w.Write([]byte(`{"result":` + string(buf) + `,"status":200}`))
}

func totalVoteByHeight(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	height := params["height"]
	h, ok := strconv.Atoi(height)
	if ok != nil || h < 0 {
		http.Error(w, `{"result":"invalid height","status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
		return
	}
	rst, err := dba.ToFloat(`select  sum(value) as value from chain_vote_info a right join chain_producer_info b on a.producer_public_key = b.ownerpublickey  where (cancel_height > ` + height + ` or cancel_height is null) and height <= ` + height + ``)
	if err != nil {
		http.Error(w, `{"result":"internal error : `+err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`{"result":` + fmt.Sprintf("%.8f", rst) + `,"status":200}`))
}

func confirmedDetailByHeight(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	height := params["height"]
	h, ok := strconv.Atoi(height)
	if ok != nil || h < 0 {
		http.Error(w, `{"result":"invalid height","status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
		return
	}
	reqBody := `{"method": "getconfirmbyheight","params": {"height":` + height + `,"verbosity":1}}`
	var resp map[string]interface{}
	var err error
	if strings.HasPrefix(config.Conf.Ela.Restful, "http") {
		resp, err = tools.PostAuth(config.Conf.Ela.Jsonrpc, reqBody, config.Conf.Ela.JsonrpcUser, config.Conf.Ela.JsonrpcPassword)
	} else {
		resp, err = tools.PostAuth("http://"+config.Conf.Ela.Jsonrpc, reqBody, config.Conf.Ela.JsonrpcUser, config.Conf.Ela.JsonrpcPassword)
	}
	if err != nil {
		http.Error(w, `{"result":"internal error : `+err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
		return
	}

	buf, err := json.Marshal(resp["result"])
	if err != nil {
		http.Error(w, `{"result":"internal error : `+err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
		return
	}
	w.Write([]byte(`{"result":` + string(buf) + `,"status":200}`))
}

func getProducerByTxs(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, `{"result":"bad request","status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
		return
	}
	data := new(map[string]interface{})
	err = json.Unmarshal(b, data)
	if err != nil {
		http.Error(w, `{"result":"bad request","status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
		return
	}
	txids, ok := (*data)["txid"].([]interface{})
	if !ok {
		http.Error(w, `{"result":"bad request","status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
		return
	}
	type ret struct {
		Producer interface{}
		Txid     string
	}
	var rst []ret
	for _, v := range txids {
		txid := v.(string)
		tmp := chain.Producer_info{}
		//TODO the transaction may contains producer that has been canceled
		producer, err := dba.ToStruct("select b.* from chain_vote_info a right join chain_producer_info b on a.producer_public_key = b.ownerpublickey where a.txid = '"+txid+"'", tmp)
		if err != nil {
			http.Error(w, `{"result":"internal error","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
			return
		}
		if len(producer) > 0 && producer[0] != nil {
			rst = append(rst, ret{
				Producer: producer,
				Txid:     txid,
			})
		}
	}
	buf, err := json.Marshal(&rst)
	if err != nil {
		http.Error(w, `{"result":"internal error : `+err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
		return
	}
	w.Write([]byte(`{"result":` + string(buf) + `,"status":200}`))
}

var version = "1.0.1"

//ping ping can be used as a heart beat
func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"result":"pong ` + version + `","status":200}`))
}

//get max height value of chain
func syncChecking(w http.ResponseWriter, r *http.Request) {
	c, err := dba.ToInt("select max(height) as height from chain_block_header")
	if err != nil {
		w.Write([]byte(`{"result":"` + err.Error() + `","status":500}`))
		return
	}
	w.Write([]byte(`{"result":` + strconv.Itoa(c) + `,"status":200}`))
}

const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="refresh" content="15">
	</head>
	<body>
		<table border='1'>
			<tr>
				<th>address</th>
				<th>txid</th>
				<th>type</th>
				<th>value</th>
				<th>createTime</th>
				<th>height</th>
				<th>fee</th>
				<th>inputs</th>
				<th>outputs</th>
				<th>memo</th>
				<th>txType</th>
			</tr>
			{{ range . }}
			<tr>
				<td>{{ .Address }}</td>
				<td>{{ .Txid }}</td>
				<td>{{ .Type }}</td>
				<td>{{ .Value }}</td>
				<td>{{ .CreateTime }}</td>
				<td>{{ .Height }}</td>
				<td>{{ .Fee }}</td>
				<td>{{ .Inputs }}</td>
				<td>{{ .Outputs }}</td>
				<td>{{ .Memo | hexToStr}}</td>
				<td>{{ .TxType }}</td>
			</tr>
			{{ end }}
		</table>
	</body>
</html>`

var (
	tplFuncMap = template.FuncMap{
		"hexToStr": func(in string) (out string) {
			b, err := hex.DecodeString(in)
			if err != nil {
				return ""
			}
			return string(b)
		},
	}
	t, _ = template.New("webpage").Funcs(tplFuncMap).Parse(tpl)
)

//list list the transaction history data
func list(w http.ResponseWriter, r *http.Request) {
	pageNum := r.FormValue("pageNum")
	txId := r.FormValue("txId")
	var sql string
	if pageNum != "" {
		pageSize := r.FormValue("pageSize")
		var size int64
		if pageSize != "" {
			var err error
			size, err = strconv.ParseInt(pageSize, 10, 64)
			if err != nil {
				w.Write([]byte(`{"result":"` + err.Error() + `","status":400}`))
				return
			}
		} else {
			size = 10
		}
		num, err := strconv.ParseInt(pageNum, 10, 64)
		if err != nil {
			w.Write([]byte(`{"result":"` + err.Error() + `","status":400}`))
			return
		}
		if num <= 0 {
			num = 1
		}
		from := (num - 1) * size
		if txId != "" {
			sql = "select * from chain_block_transaction_history where txType = 'TransferAsset' and txid = '" + txId + "' order by id desc limit " + strconv.FormatInt(from, 10) + "," + strconv.FormatInt(size, 10)
		} else {
			sql = "select * from chain_block_transaction_history where txType = 'TransferAsset' order by id desc limit " + strconv.FormatInt(from, 10) + "," + strconv.FormatInt(size, 10)
		}
	} else {
		if txId != "" {
			sql = "select * from chain_block_transaction_history where txType = 'TransferAsset' and txid = '" + txId + "' order by id desc limit 100"
		} else {
			sql = "select * from chain_block_transaction_history where txType = 'TransferAsset' order by id desc limit 100"
		}
	}

	list, err := dba.ToStruct(sql, chain.Block_transaction_history{})

	if err != nil {
		w.Write([]byte(`{"result":"` + err.Error() + `","status":500}`))
		return
	}

	err = t.Execute(w, list)

	if err != nil {
		w.Write([]byte(`{"result":"` + err.Error() + `","status":500}`))
		return
	}
}

//getBtcBlockHeight get bitcoin current blockchain height
func getBtcBlockHeight(w http.ResponseWriter, r *http.Request) {
	helper := rpchelper{}
	height, err, status := helper.getBestheight()
	if err != nil {
		w.Write([]byte(`{"result":"` + err.Error() + `","status":` + strconv.Itoa(status) + `}`))
		return
	}
	w.Write([]byte(`{"result":` + strconv.Itoa(int(height)) + `,"status":200}`))
}

//getBtcTransaction get bitcoin transaction
func getBtcTransaction(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	helper := rpchelper{param}
	transaction, err, status := helper.getTransaction()
	if err != nil {
		w.Write([]byte(`{"result":"` + err.Error() + `","status":` + strconv.Itoa(status) + `}`))
		return
	}
	w.Write([]byte(`{"result":` + transaction + `,"status":200}`))
}

//getBtcBalance get bitcoin balance of the requested address
func getBtcBalance(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	helper := rpchelper{param}
	balance, err, status := helper.getBalance()
	if err != nil {
		w.Write([]byte(`{"result":"` + err.Error() + `","status":` + strconv.Itoa(status) + `}`))
		return
	}
	w.Write([]byte(`{"result":` + strconv.FormatFloat(balance, 'f', 8, 64) + `,"status":200}`))
}

//getBtcBlock get bitcoin block info
func getBtcBlock(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	helper := rpchelper{param}
	block, err, status := helper.getBlockDetail()
	if err != nil {
		w.Write([]byte(`{"result":"` + err.Error() + `","status":` + strconv.Itoa(status) + `}`))
		return
	}
	w.Write([]byte(`{"result":` + block + `,"status":200}`))

}

//getCmcPrice get price from cmc
func getCmcPrice(w http.ResponseWriter, r *http.Request) {
	limit := r.FormValue("limit")
	if limit == "" {
		http.Error(w, `{"result":"invalid request param : limit can not be blank" ,"status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
		return
	}
	ilimit, err := strconv.Atoi(limit)
	if err != nil {
		http.Error(w, `{"result":"invalid request param : limit must be a number " ,"status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
		return
	}
	if ilimit > config.Conf.Cmc.NumOfCoin {
		http.Error(w, `{"result":"invalid request param : limit exceed maximum value " ,"status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
		return
	}
	_id, err := dba.ToInt("select _id from chain_cmc_price where symbol = 'BTC' order by _id desc limit 1")
	if err != nil {
		http.Error(w, `{"result":"internal error : `+err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
		return
	}
	l, err := dba.Query("select * from chain_cmc_price where _id between " + strconv.Itoa(_id) + " and " + strconv.Itoa(ilimit+_id-1))
	if err != nil {
		http.Error(w, `{"result":"internal error : `+err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
		return
	}
	ret := [][]byte{
		[]byte("["),
	}
	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		m := e.Value.(map[string]interface{})
		buf, err := json.Marshal(m)
		if err != nil {
			http.Error(w, `{"result":"internal error : `+err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
			return
		}
		ret = append(ret, buf)
		if i != l.Len()-1 {
			ret = append(ret, []byte(","))
		}
		i++
	}
	ret = append(ret, []byte("]"))
	w.Write(bytes.Join(ret, nil))
}

func postRpc(w http.ResponseWriter, r *http.Request) {
	if config.Conf.Eth.Enable {
		b, err := ioutil.ReadAll(r.Body)
		data, err := tools.Post(config.Conf.Eth.Endpoint, string(b))
		if err != nil {
			http.Error(w, `{"result":"internal error : `+err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
			return
		}
		ret, err := json.Marshal(data)
		if err != nil {
			http.Error(w, `{"result":"internal error : `+err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
			return
		}
		w.Write(ret)
	} else {
		http.Error(w, `{"result":" Eth service is not enabled  ","status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
	}
}

func getEthHistory(w http.ResponseWriter, r *http.Request) {
	if config.Conf.Eth.Enable {
		var account string
		var err error
		if strings.ToUpper(r.Method) == "POST"  {
			b, err := ioutil.ReadAll(r.Body)
			var req map[string]string
			err = json.Unmarshal(b, &req)
			if err != nil {
				http.Error(w, `{"result":"invalid request : `+err.Error()+`","status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
				return
			}
			account = req["account"]
			if account == "" {
				http.Error(w, `{"result":"invalid request : `+err.Error()+`","status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
				return
			}
		}else {
			account = r.FormValue("address")
		}
		if account == "" {
			http.Error(w, `{"result":"invalid request : `+err.Error()+`","status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
			return
		}
		history, err := chain.GetEthHistory(account)
		if err != nil {
			http.Error(w, `{"result":"invalid request : `+err.Error()+`","status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
			return
		}
		resp := make(map[string]interface{})
		resp["status"] = 1
		resp["message"] = "OK"
		resp["result"] = history
		retBuf, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, `{"result":"internal error : `+err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
			return
		}
		w.Write(retBuf)
	} else {
		http.Error(w, `{"result":" Eth service is not enabled  ","status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
	}

}
