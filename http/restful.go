package http

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/elastos/Elastos.ORG.API.Misc/chain"
	"github.com/elastos/Elastos.ORG.API.Misc/config"
	"github.com/elastos/Elastos.ORG.API.Misc/db"
	"github.com/elastos/Elastos.ORG.API.Misc/tools"
	"github.com/engoengine/math"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	dba    		= db.NewInstance()
)

func StartServer() {
	http.ListenAndServe(":"+config.Conf.ServerPort, router)
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
			w.Write([]byte(`{"result":"`+ err.Error() + `","status":500}`))
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

func producerStatistic(w http.ResponseWriter,r *http.Request){
	params := mux.Vars(r)
	pub := params["producer"]
	if pub == "" || len(pub) != 66 {
		http.Error(w,`{"result":"invalid public key","status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
		return
	}
	rst , err := dba.ToStruct("select * from chain_vote_info where producer_public_key = '"+pub+"' and (outputlock = 0 or outputlock >= height) and is_valid = 'YES'",chain.Vote_info{})
	if err != nil {
		http.Error(w,`{"result":"internal error : `+ err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
		return
	}
	buf , err := json.Marshal(&rst)
	if err != nil {
		http.Error(w,`{"result":"internal error : `+ err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
		return
	}
	w.Write([]byte(`{"result":` + string(buf) + `,"status":200}`))
}

func voterStatistic(w http.ResponseWriter,r *http.Request){
	params := mux.Vars(r)
	addr := params["address"]
	if addr == "" || len(addr) != 34 {
		http.Error(w,`{"result":"invalid address","status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
		return
	}
	rst , err := dba.ToStruct("select * from chain_vote_info where address = '"+addr+"' and (outputlock = 0 or outputlock >= height) and is_valid = 'YES'",chain.Vote_info{})
	if err != nil {
		http.Error(w,`{"result":"internal error : `+ err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
		return
	}
	buf , err := json.Marshal(&rst)
	if err != nil {
		http.Error(w,`{"result":"internal error : `+ err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
		return
	}
	w.Write([]byte(`{"result":` + string(buf) + `,"status":200}`))
}

var version = "1.0.1"

//ping ping can be used as a heart beat
func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"result":"pong `+version+`","status":200}`))
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
			sql = "select * from chain_block_transaction_history where txType = 'TransferAsset' and txid = '"+txId+"' order by id desc limit " + strconv.FormatInt(from, 10) + "," + strconv.FormatInt(size, 10)
		}else {
			sql = "select * from chain_block_transaction_history where txType = 'TransferAsset' order by id desc limit " + strconv.FormatInt(from, 10) + "," + strconv.FormatInt(size, 10)
		}
	} else {
		if txId != "" {
			sql = "select * from chain_block_transaction_history where txType = 'TransferAsset' and txid = '"+txId+"' order by id desc limit 100"
		}else {
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
	height , err , status := helper.getBestheight()
	if err != nil {
		w.Write([]byte(`{"result":"` + err.Error() + `","status":`+strconv.Itoa(status)+`}`))
		return
	}
	w.Write([]byte(`{"result":` + strconv.Itoa(int(height)) + `,"status":200}`))
}

//getBtcTransaction get bitcoin transaction
func getBtcTransaction(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	helper := rpchelper{param}
	transaction , err , status := helper.getTransaction()
	if err != nil {
		w.Write([]byte(`{"result":"` + err.Error() + `","status":`+strconv.Itoa(status)+`}`))
		return
	}
	w.Write([]byte(`{"result":` + transaction + `,"status":200}`))
}

//getBtcBalance get bitcoin balance of the requested address
func getBtcBalance(w http.ResponseWriter,r *http.Request) {
	param := mux.Vars(r)
	helper := rpchelper{param}
	balance , err , status := helper.getBalance()
	if err != nil {
		w.Write([]byte(`{"result":"` + err.Error() + `","status":`+strconv.Itoa(status)+`}`))
		return
	}
	w.Write([]byte(`{"result":` + strconv.FormatFloat(balance,'f',8,64) + `,"status":200}`))
}

//getBtcBlock get bitcoin block info
func getBtcBlock(w http.ResponseWriter,r *http.Request) {
	param := mux.Vars(r)
	helper := rpchelper{param}
	block , err , status := helper.getBlockDetail()
	if err != nil {
		w.Write([]byte(`{"result":"` + err.Error() + `","status":`+strconv.Itoa(status)+`}`))
		return
	}
	w.Write([]byte(`{"result":` + block + `,"status":200}`))

}

//getCmcPrice get price from cmc
func getCmcPrice(w http.ResponseWriter,r *http.Request){
	apiKey := r.Header["Apikey"]
	tp_param := r.Header["Timestamp"]
	if len(apiKey) == 0 || len(tp_param) == 0 {
		http.Error(w, `{"result":"invalid request param : apiKey or timestamp can not be blank" ,"status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
		return
	}
	itp_param , err := strconv.ParseInt(tp_param[0],10,64)
	if err != nil {
		http.Error(w, `{"result":"invalid request param : invalid timestamp","status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
		return
	}
	tp_local := time.Now().UTC().Unix() * 1000
	if math.Abs(float32(tp_local - itp_param))/(1000 * 60) > 5 {
		http.Error(w, `{"result":"invalid request param : apiKey out of date","status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
		return
	}
	keyHash := sha256.Sum256([]byte(config.Conf.VisitKey+tp_param[0]))
	if hex.EncodeToString(keyHash[:]) != apiKey[0] {
		http.Error(w, `{"result":"invalid request param : apiKey not correct ","status":`+strconv.Itoa(http.StatusBadRequest)+`}`, http.StatusBadRequest)
		return
	}
	limit := r.FormValue("limit")
	if limit == "" {
		http.Error(w, `{"result":"invalid request param : limit can not be blank" ,"status":`+strconv.Itoa(http.StatusBadRequest)+`}` , http.StatusBadRequest)
		return
	}
	ilimit , err := strconv.Atoi(limit)
	if err != nil {
		http.Error(w, `{"result":"invalid request param : limit must be a number " ,"status":`+strconv.Itoa(http.StatusBadRequest)+`}` , http.StatusBadRequest)
		return
	}
	if ilimit > config.Conf.Cmc.NumOfCoin {
		http.Error(w, `{"result":"invalid request param : limit exceed maximum value " ,"status":`+strconv.Itoa(http.StatusBadRequest)+`}` , http.StatusBadRequest)
		return
	}
	_id , err := dba.ToInt("select _id from chain_cmc_price where symbol = 'BTC' order by _id desc limit 1")
	if err != nil {
		http.Error(w,`{"result":"internal error : `+ err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
		return
	}
	l , err := dba.Query("select * from chain_cmc_price where _id between " +strconv.Itoa(_id) + " and " + strconv.Itoa(ilimit + _id - 1))
	if err != nil {
		http.Error(w,`{"result":"internal error : `+ err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
		return
	}
	ret := [][]byte{
		[]byte("["),
	}
	i := 0
	for e := l.Front() ; e !=nil ;e = e.Next(){
		m := e.Value.(map[string]interface{})
		buf , err := json.Marshal(m)
		if err != nil {
			http.Error(w,`{"result":"internal error : `+ err.Error()+`","status":`+strconv.Itoa(http.StatusInternalServerError)+`}`, http.StatusInternalServerError)
			return
		}
		ret = append(ret,buf)
		if i != l.Len() - 1 {
			ret = append(ret,[]byte(","))
		}
		i++
	}
	ret = append(ret,[]byte("]"))
	w.Write(bytes.Join(ret,nil))
}
