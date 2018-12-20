package http

import (
	"encoding/hex"
	"encoding/json"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/elastos/Elastos.ORG.API.Misc/chain"
	"github.com/elastos/Elastos.ORG.API.Misc/config"
	"github.com/elastos/Elastos.ORG.API.Misc/db"
	"github.com/elastos/Elastos.ORG.API.Misc/log"
	"github.com/elastos/Elastos.ORG.API.Misc/tools"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var (
	dba    = db.NewInstance()
	client *rpcclient.Client
)

func StartServer() {
	http.ListenAndServe(":"+config.Conf.ServerPort, router)
}

func init() {
	if config.Conf.Btc.Host != "" {
		go func() {
			var err error
			client, err = rpcclient.New(&rpcclient.ConnConfig{
				HTTPPostMode: true,
				DisableTLS:   true,
				Host:         config.Conf.Btc.Host,
				User:         config.Conf.Btc.Rpcuser,
				Pass:         config.Conf.Btc.Rpcpasswd,
			}, nil)
			if err != nil {
				log.Error("Error Connect to Bitcoin node :" , err.Error())
			}
		}()
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
		sql = "select txid,type,value,createTime,height,inputs,outputs,fee,txType,memo from chain_block_transaction_history where address = '" + address + "' limit " + strconv.FormatInt(from, 10) + "," + strconv.FormatInt(size, 10)
	} else {
		sql = "select txid,type,value,createTime,height,inputs,outputs,fee,txType,memo from chain_block_transaction_history where address = '" + address + "'"
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

//ping ping can be used as a heart beat
func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"result":"pong","status":200}`))
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
		sql = "select * from chain_block_transaction_history where txType = 'TransferAsset' order by id desc limit " + strconv.FormatInt(from, 10) + "," + strconv.FormatInt(size, 10)
	} else {
		sql = "select * from chain_block_transaction_history where txType = 'TransferAsset' order by id desc limit 100"
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
	blockHeight , err := client.GetBlockCount()
	if err != nil {
		w.Write([]byte(`{"result":"` + err.Error() + `","status":500}`))
		return
	}
	w.Write([]byte(`{"result":` + strconv.Itoa(int(blockHeight)) + `,"status":200}`))
}

//getBtcTransaction get bitcoin transaction
func getBtcTransaction(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	txid := param["txid"]
	btxid , err := hex.DecodeString(txid)
	if err != nil || len(btxid) != 32{
		w.Write([]byte(`{"result":"Invalid txid","status":400}`))
		return
	}

	hash := &chainhash.Hash{}
	hash.SetBytes(tools.ReverseBytes(btxid))
	tx , err := client.GetRawTransactionVerbose(hash)

	if err != nil {
		w.Write([]byte(`{"result":"` + err.Error() + `","status":500}`))
		return
	}
	buf , _ :=json.Marshal(tx)
	w.Write([]byte(`{"result":` + string(buf) + `,"status":200}`))
}
