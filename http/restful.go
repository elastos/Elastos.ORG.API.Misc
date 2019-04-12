package http

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/elastos/Elastos.ELA.Utility/crypto"
	"github.com/elastos/Elastos.ORG.API.Misc/chain"
	"github.com/elastos/Elastos.ORG.API.Misc/config"
	"github.com/elastos/Elastos.ORG.API.Misc/db"
	"github.com/elastos/Elastos.ORG.API.Misc/log"
	"github.com/elastos/Elastos.ORG.API.Misc/tools"
	"github.com/gorilla/mux"
	"html/template"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	dba = db.NewInstance()
)

func StartServer() {
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
		sql = "select address,txid,type,value,createTime,height,inputs,outputs,fee,txType,memo from chain_block_transaction_history where address = '" + address + "' limit " + strconv.FormatInt(from, 10) + "," + strconv.FormatInt(size, 10)
	} else {
		sql = "select address,txid,type,value,createTime,height,inputs,outputs,fee,txType,memo from chain_block_transaction_history where address = '" + address + "'"
	}
	l, err := dba.Query(sql)
	if err != nil {
		w.Write([]byte(`{"result":"` + err.Error() + `","status":500}`))
		return
	}
	bhs := make([]chain.Block_transaction_history, 0)
	totalNum := 0
	for e := l.Front(); e != nil; e = e.Next() {
		history := new(chain.Block_transaction_history)
		line := e.Value.(map[string]interface{})
		tools.Map2Struct(line, history)
		inputsArr := strings.Split(line["inputs"].(string), ",")
		history.Inputs = inputsArr[:len(inputsArr)-1]
		outputsArr := strings.Split(line["outputs"].(string), ",")
		history.Outputs = outputsArr[:len(outputsArr)-1]
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
	l, err = dba.Query("select count(*) as count from chain_block_transaction_history where address = '" + address + "'")
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

func producerStatistic(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pub := params["producer"]
	if pub == "" || len(pub) != 66 {
		http.Error(w, `{"result":"invalid public key","status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
		return
	}
	rst, err := dba.ToStruct("select * from chain_vote_info where producer_public_key = '"+pub+"' and (outputlock = 0 or outputlock >= height) and is_valid = 'YES'", chain.Vote_info{})
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
	rst, err := dba.ToStruct("select * from chain_vote_info where address = '"+addr+"' and (outputlock = 0 or outputlock >= height) and is_valid = 'YES'", chain.Vote_info{})
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

func producerRankByHeight(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	height := params["height"]
	h, ok := strconv.Atoi(height)
	if ok != nil || h < 0 {
		http.Error(w, `{"result":"invalid height","status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
		return
	}
	rst, err := dba.ToStruct(`select a.* , (@row_number:=@row_number + 1) as "rank",b.* from 
(select A.producer_public_key , sum(value) as value from chain_vote_info A where A.cancel_height > `+height+` or
 cancel_height is null group by producer_public_key order by value desc) a inner join chain_producer_info b on a.producer_public_key = b.ownerpublickey 
 ,  (SELECT @row_number:=0) AS t`, chain.Vote_info{})
	if err != nil {
		http.Error(w, `{"result":"internal error : `+err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
		return
	}
	roundStartHeight := h - (h-config.Conf.DposStartHeight)%36
	roundStartHeightTotalVote, err := dba.ToFloat(`select sum(value) as value from chain.chain_vote_info where cancel_height > ` + strconv.Itoa(roundStartHeight) + ` or cancel_height is null `)
	if err != nil {
		http.Error(w, `{"result":"internal error : `+err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
		return
	}
	for _, r := range rst {
		vi := r.(*chain.Vote_info)
		addr, err := getAddress(vi.Ownerpublickey)
		if err != nil {
			log.Warn("Invalid Ownerpublickey " + vi.Ownerpublickey)
			continue
		}
		vi.Address = addr
		val, err := dba.ToString("select value from chain_block_transaction_history where height = " + height + " and txType = 'CoinBase' and value < " + strconv.Itoa(tools.Miner_Reward_PerBlock) + " and address = '" + addr + "'")
		if err != nil {
			log.Warn("Invalid Ownerpublickey " + vi.Ownerpublickey)
			continue
		}
		if val != "" {
			vi.Reward = val
		}
		vote, err := strconv.ParseFloat(vi.Value, 64)
		if err != nil {
			http.Error(w, `{"result":"internal error : `+err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
			return
		}
		vi.EstRewardPerYear = strconv.FormatFloat(float64(175834088/(roundStartHeightTotalVote*100000000)*vote*365*720), 'f', 8, 64)
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
	rst, err := dba.ToFloat(`select  sum(value) as value from chain.chain_vote_info where cancel_height > ` + height + ` or cancel_height is null `)
	if err != nil {
		http.Error(w, `{"result":"internal error : `+err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`{"result":` + fmt.Sprintf("%.8f", rst) + `,"status":200}`))
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
	apiKey := r.Header["Apikey"]
	tp_param := r.Header["Timestamp"]
	if len(apiKey) == 0 || len(tp_param) == 0 {
		http.Error(w, `{"result":"invalid request param : apiKey or timestamp can not be blank" ,"status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
		return
	}
	itp_param, err := strconv.ParseInt(tp_param[0], 10, 64)
	if err != nil {
		http.Error(w, `{"result":"invalid request param : invalid timestamp","status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
		return
	}
	tp_local := time.Now().UTC().Unix() * 1000
	if math.Abs(float64(tp_local-itp_param))/(1000*60) > 5 {
		http.Error(w, `{"result":"invalid request param : apiKey out of date","status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
		return
	}
	keyHash := sha256.Sum256([]byte(config.Conf.VisitKey + tp_param[0]))
	if hex.EncodeToString(keyHash[:]) != apiKey[0] {
		http.Error(w, `{"result":"invalid request param : apiKey not correct ","status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
		return
	}
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

func getAddress(publicKeyHex string) (string, error) {
	publicKey, err := hex.DecodeString(publicKeyHex)
	if err != nil {
		return "", err
	}
	pub, err := crypto.DecodePoint(publicKey)
	if err != nil {
		return "", err
	}
	code, err := crypto.CreateStandardRedeemScript(pub)
	if err != nil {
		return "", err
	}
	hash, err := crypto.ToProgramHash(code)
	if err != nil {
		return "", err
	}
	addr, err := hash.ToAddress()
	if err != nil {
		return "", err
	}
	return addr, nil
}
