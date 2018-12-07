package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	router *mux.Router
	db *sql.DB
	Conf = Config{}
)

type Config struct {
	Node , ServerPort string
}

type Address_history struct{
	History []Block_transaction_history
	TotalNum int
}

type Block_transaction_history struct {
	Txid string
	Type string
	Value int64
	CreateTime int64
	Height int
	Fee int64
	Inputs  []string
	Outputs []string
}

const (
	BbName = "chain.db"
	BlockHeight = "/api/v1/block/height"
	BlockDetail = "/api/v1/block/details/height/"
	TransactionDetail = "/api/v1/transaction/"
	INCOME = "income"
	SPEND = "spend"
	ERROR_REQUEST = "Error Request :"
	ELA = 100000000
	MINEING_ADDR = "0000000000000000000000000000000000"
	CROSS_CHAIN_FEE = 10000
)

func main(){
	go SyncChain()
	defer db.Close()
	http.ListenAndServe(":"+Conf.ServerPort,router)
}


func SyncChain() {
	for {
		tx , err := db.Begin()
		if err = sync(tx) ; err != nil {
			log.Printf("Sync Height Error : %v \n" ,  err.Error())
			tx.Rollback()
		}else{
			tx.Commit()
		}
		<-time.After(time.Second * 10)
	}
}

func sync(tx *sql.Tx) error{

	resp , err := Get(Conf.Node + BlockHeight)

	if err != nil {
		return err
	}

	r , err := tx.Query("select height from block_CurrHeight order by id desc limit 1")
	if err != nil {
		return err
	}
	defer r.Close()
	storeHeight :=-1
	if r.Next() {
		r.Scan(&storeHeight)
	}

	chainHeight , ok := resp["Result"]

	if ok {
		if storeHeight == int(chainHeight.(float64)) {
			return nil
		}

		count := 0
		isSet := false
		for curr := storeHeight + 1 ; curr <= int(chainHeight.(float64)) ; curr++{
			count++
			resp , err = Get(Conf.Node + BlockDetail + strconv.FormatInt(int64(curr),10))
			txArr := (resp["Result"].(map[string]interface{}))["tx"].([]interface{})
			if len(txArr) == 0 {
				continue
			}
			for _ , v := range txArr {
				stmt , err := tx.Prepare("insert into block_transaction_history (address,txid,type,value,createTime,height,fee,inputs,outputs) values(?,?,?,?,?,?,?,?,?)")
				if err != nil {
					return err
				}
				defer stmt.Close()
				vm := v.(map[string]interface{})
				txid := vm["txid"].(string)
				time := vm["blocktime"].(float64)
				t := vm["type"].(float64)
				_type := INCOME
				if t == 0 {
					vout :=  vm["vout"].([]interface{})
					coinbase := make([]map[string]interface{},0)
					to := ""
					for _ , vv := range vout{
						vvm := vv.(map[string]interface{})
						value := vvm["value"].(string)
						address := vvm["address"].(string)
						coinbaseMap := make(map[string]interface{})
						fv , err := strconv.ParseFloat(value,64)
						if err != nil {
							return err
						}
						coinbaseMap["value"] = int64(fv * ELA)
						coinbaseMap["address"] = address
						coinbase = append(coinbase,coinbaseMap)
						to += address +","

					}

					for _ , v := range coinbase {
						_ , err := stmt.Exec(v["address"],txid,_type,v["value"],strconv.FormatFloat(time,'f',0,64),curr,0,MINEING_ADDR,to)
						if err != nil {
							return err
						}
					}

				}else{

					vin :=  vm["vin"].([]interface{})
					spend := make(map[string]float64)
					totalInput := 0.0
					from := ""
					to := ""
					for _ , vv := range vin{
						vvm := vv.(map[string]interface{})
						vintxid := vvm["txid"].(string)
						vinindex := vvm["vout"].(float64)
						txResp , err := Get(Conf.Node + TransactionDetail+vintxid)
						if err != nil {
							return err
						}
						vout := ((txResp["Result"].(map[string]interface{}))["vout"].([]interface{}))[int(vinindex)]
						voutm := vout.(map[string]interface{})
						address := voutm["address"].(string)
						value, err:= strconv.ParseFloat(voutm["value"].(string),64)
						totalInput += value
						if err != nil {
							return err
						}
						v , ok := spend[address]
						if ok {
							spend[address] = v + value
						}else {
							spend[address] = value
						}
						from += address + ","
					}
					vout :=  vm["vout"].([]interface{})
					receive := make(map[string]float64)
					totalOutput := 0.0
					for _ , vv := range vout{
						vvm := vv.(map[string]interface{})
						value, err:= strconv.ParseFloat( vvm["value"].(string),64)
						totalOutput += value
						if err != nil {
							return err
						}
						address := vvm["address"].(string)
						v , ok := receive[address]
						if ok {
							receive[address] = v + value
						}else {
							receive[address] = value
						}
						to += address +","
					}
					fee := int64(math.Round((totalInput - totalOutput)*ELA))
					if fee < 0 {
						fee = CROSS_CHAIN_FEE
					}
					for k , r := range receive {
						_type = INCOME
						s , ok := spend[k]
						var value float64
						if ok {
							if s > r {
								value = math.Round((s - r)*ELA)
								_type = SPEND
							}else {
								value = math.Round((r - s)*ELA)
							}
							delete(spend,k)
						}else {
							value = math.Round(r*ELA)
						}
						_ , err := stmt.Exec(k,txid,_type,int64(value),strconv.FormatFloat(time,'f',0,64),curr,fee,from,to)
						if err != nil {
							return err
						}
					}

					for k , r := range spend {
						_type = SPEND
						_ , err := stmt.Exec(k,txid,_type,int64(r*ELA),strconv.FormatFloat(time,'f',0,64),curr,fee,from,to)
						if err != nil {
							return err
						}
					}
				}
			}
			if count % 5000 == 0 {
				isSet = true
				s , err := tx.Prepare("insert into block_CurrHeight (height) values(?)")
				if err != nil {
					return err
				}
				defer s.Close()
				_ , err = s.Exec(curr)
				if err != nil {
					return err
				}
				return nil
			}
		}
		if !isSet {
			s , err := tx.Prepare("insert into block_CurrHeight (height) values(?)")
			if err != nil {
				return err
			}
			defer s.Close()
			_ , err = s.Exec(chainHeight)
			if err != nil {
				return err
			}
		}
	}

	return nil
}


func init(){
	router = mux.NewRouter()
	router.HandleFunc("/api/1/history/{addr}",history).Methods("GET")
	initDb()
	initConfig()
}


func history(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	address := params["addr"]
	pageNum := r.FormValue("pageNum")
	var sql string
	if pageNum != "" {
		pageSize := r.FormValue("pageSize")
		var size int64
		if pageSize != "" {
			var err error
			size , err = strconv.ParseInt(pageSize ,10,64)
			if err != nil {
				w.Write([]byte(`{"result":"`+err.Error()+`",status:400}`))
				return
			}
		}else{
			size = 10
		}
		num , err := strconv.ParseInt(pageNum ,10,64)
		if err != nil {
			w.Write([]byte(`{"result":"`+err.Error()+`",status:400}`))
			return
		}
		if num <= 0 {
			num = 1
		}
		from := (num -1) * size
		sql = "select txid,type,value,createTime,height,inputs,outputs,fee from block_transaction_history where address = ? limit " + strconv.FormatInt(from,10) + "," + strconv.FormatInt(size,10)
	}else{
		sql = "select txid,type,value,createTime,height,inputs,outputs,fee from block_transaction_history where address = ?"
	}
	s , err := db.Prepare(sql)
	if err != nil{
		w.Write([]byte(`{"result":"`+ERROR_REQUEST + err.Error()+`",status:500}`))
		return
	}
	rst , err := s.Query(address)
	if err != nil{
		w.Write([]byte(`{"result":"`+ERROR_REQUEST + err.Error()+`",status:500}`))
		return
	}
	bhs := make([]Block_transaction_history,0)
	totalNum := 0
	for rst.Next() {
		history :=Block_transaction_history{}
		var inputs , outputs string
		err := rst.Scan(&history.Txid,&history.Type,&history.Value,&history.CreateTime,&history.Height,&inputs,&outputs,&history.Fee)
		inputsArr := strings.Split(inputs,",")
		history.Inputs = inputsArr[:len(inputsArr)-1]
		outputsArr := strings.Split(outputs,",")
		history.Outputs = outputsArr[:len(outputsArr)-1]
		if err != nil {
			w.Write([]byte(`{"result":"`+ERROR_REQUEST + err.Error()+`",status:500}`))
			return
		}
		bhs = append(bhs,history)
	}
	defer func() {s.Close() ; rst.Close()}()
	ss , err := db.Prepare("select count(*) as count from block_transaction_history where address = ?")
	if err != nil {
		w.Write([]byte(`{"result":"`+ERROR_REQUEST + err.Error()+`",status:500}`))
		return
	}
	rsts , err := ss.Query(address)
	if err != nil {
		w.Write([]byte(`{"result":"`+ERROR_REQUEST + err.Error()+`",status:500}`))
		return
	}
	rsts.Next()
	rsts.Scan(&totalNum)
	defer func() {ss.Close() ; rsts.Close()}()
	addrHis := Address_history{bhs,totalNum}
	buf , err := json.Marshal(&addrHis)
	if err != nil {
		w.Write([]byte(`{"result":"`+ERROR_REQUEST + err.Error()+`",status:500}`))
		return
	}
	w.Write([]byte(`{"result":` + string(buf)+`,"status":200}`))
}

func readBody(r *http.Request) (string){
	body , err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return string(body)
}

func Get(url string) (map[string]interface{},error) {
	log.Printf("Request URL = %v \n", url)
	r , err := http.Get(url)
	if err != nil {
		return nil , err
	}
	resp , err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil , err
	}
	rstMap := make(map[string]interface{})
	json.Unmarshal(resp,&rstMap)
	return rstMap , nil
}


func initDb(){
	var err error
	db, err = sql.Open("sqlite3", BbName)
	if err != nil {
		log.Fatal(err)
	}
	initializeTable()
}

func initializeTable(){
	createTableSqlStmtArr := []string{
		`create table IF not exists block_currHeight (id integer not null primary key , height integer not null);`,
		`create table IF not exists block_transaction_history (id integer not null primary key , address varchar(34) not null ,
		txid varchar(64) not null ,type blob not null, value integer not null, createTime integer not null , height integer not null,fee integer not null,
		inputs blob  not null ,outputs blob not null);`,
	}

	for _ , v := range createTableSqlStmtArr {
		log.Printf("Execute sql :%v",v)
		_, err := db.Exec(v)
		if err != nil {
			log.Printf("Error execute sql : %q \n", err, v)
			return
		}
	}

}

func initConfig(){
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&Conf)
	if err != nil {
		log.Fatal("Error init Config :", err)
	}
}
