package db

import (
	"container/list"
	"database/sql"
	"errors"
	"github.com/elastos/Elastos.ORG.API.Misc/config"
	"github.com/elastos/Elastos.ORG.API.Misc/log"
	"github.com/elastos/Elastos.ORG.API.Misc/tools"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"strconv"
	"time"
)

type Dialect struct {
	db              *sql.DB
	driverSource    string
	maxConns        int
	maxIdles        int
	connMaxLifeTime time.Duration

}

func NewInstance() (*Dialect){
	var dia = new(Dialect)
	err := dia.Create(config.Conf.DbDriverSource)
	if err != nil {
		log.Fatalf("init Db Error : %s ",err.Error())
	}
	return dia
}

//Create create a db instance .
func (dia *Dialect) Create(driverSource string) error {
	if driverSource == "" {
		return errors.New("driver source can not be blank")
	}
	db, err := sql.Open("mysql",
		driverSource)
	if err != nil {
		return err
	}
	if dia.maxConns == 0 {
		db.SetMaxOpenConns(10)
	} else {
		db.SetMaxOpenConns(dia.maxConns)
	}
	if dia.maxIdles == 0 {
		db.SetMaxIdleConns(5)
	} else {
		db.SetMaxIdleConns(dia.maxIdles)
	}
	if dia.connMaxLifeTime == 0 {
		db.SetConnMaxLifetime(time.Second * 14440)
	} else {
		db.SetConnMaxLifetime(dia.connMaxLifeTime)
	}
	dia.driverSource = driverSource
	dia.db = db
	return nil
}

//isConnected check if a db connection is still alive
func (dia *Dialect) isConnected() error {
	error := dia.db.Ping()
	if error != nil {
		return error
	}
	return nil
}

//Begin begin a transaction
func (dia *Dialect) Begin() (tx *sql.Tx, err error) {
	tx, err = dia.db.Begin()
	return
}

//Commit commit a transaction
func (dia *Dialect) Commit(tx *sql.Tx) error {
	return tx.Commit()
}

//Rollback rollback a transaction
func (dia *Dialect) Rollback(tx *sql.Tx) error {
	return tx.Rollback()
}

//Execute data manipulate language
func (dia *Dialect) Execute(sql string) (int64, error) {
	log.Info(sql)
	stmt, err := dia.db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec()
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if id == 0 {
		id, _ = result.RowsAffected()
		if err != nil {
			return 0, err
		}
	}
	return id, nil
}

//BatchExecute batch data manipulate
func (dia *Dialect) BatchExecute(sql string, tx *sql.Tx) (int64, error) {
	log.Info(sql)
	stmt, err := tx.Prepare(sql)
	if err != nil {
		return 0, err
	}
	result, err := stmt.Exec()
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if id == 0 {
		id, _ = result.RowsAffected()
		if err != nil {
			return 0, err
		}
	}
	return id, nil
}

func (dia *Dialect) Query(s string) (*list.List, error) {
	log.Info(s)
	rows, err := dia.db.Query(s)
	if err != nil {
		return nil, err
	}
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	values := make([][]byte, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	retList := list.New()

	// Fetch rows
	for rows.Next() {
		retMap := make(map[string]interface{})
		retList.PushBack(retMap)
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}

		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				// skip nil value
				continue
			}
			retMap[columns[i]] = string(col)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return retList, nil
}

func (dia *Dialect) ToStruct(sql string,strct interface{}) ([]interface{},error){
	l , err := dia.Query(sql)
	if err != nil {
		return nil,err
	}
	i := 0
	r := make([]interface{},l.Len())
	for e := l.Front(); e != nil ; e = e.Next() {
		v := reflect.New(reflect.TypeOf(strct))
		src := e.Value.(map[string]interface{})
		vi := v.Interface()
		tools.Map2Struct(src,vi)
		r[i] = reflect.ValueOf(vi).Interface()
		i++
	}
	return r,nil
}

func (dia *Dialect) ToInt(sql string) (int , error){

	l , err := dia.Query(sql)
	if err != nil {
		return -1 , err
	}
	m := l.Front().Value.(map[string]interface{})
	for _ , v := range m {
		return strconv.Atoi(v.(string))
	}

	return  -1 , err
}

func (dia *Dialect) ToString(sql string) (string , error){

	l , err := dia.Query(sql)
	if err != nil {
		return "" , err
	}
	m := l.Front().Value.(map[string]interface{})
	for _ , v := range m {
		return v.(string),nil
	}

	return  "" , err
}


func (dia *Dialect) Close() error {
	return dia.db.Close()
}

func (dia *Dialect) SetMaxOpenConnections(maxConn int) error {
	if maxConn <= 0 {
		return errors.New("maxConn must bigger than 0")
	}
	dia.maxConns = maxConn
	return nil
}

func (dia *Dialect) SetMaxIdles(maxIdles int) error {
	if maxIdles <= 0 {
		return errors.New("maxIdles must bigger than 0")
	}
	dia.maxIdles = maxIdles
	return nil
}

func (dia *Dialect) SetConnMaxLifeTime(connMaxLifeTime int) error {
	if connMaxLifeTime <= 0 {
		return errors.New("connMaxLifeTime must bigger than 0")
	}
	dia.maxIdles = connMaxLifeTime
	return nil
}