package db

import (
	"fmt"
	"github.com/elastos/Elastos.ORG.API.Misc/log"
	"github.com/stretchr/testify/require"
	"testing"
)

var driverSource = "root:12345678@tcp(127.0.0.1:3306)/chain"

func Test_Create(t *testing.T){

	should := require.New(t)
	dialect := new(Dialect)
	defer dialect.Close()
	err := dialect.Create(driverSource)
	should.NoError(err,"Error creating a dialect")

}

func Test_DML(t *testing.T){

	should := require.New(t)
	dialect := new(Dialect)
	defer dialect.Close()
	err := dialect.Create(driverSource)
	should.NoError(err)
	_ , err = dialect.Execute("delete from test")
	should.NoError(err)

	tx , err := dialect.Begin()
	should.NoError(err)
	dialect.BatchExecute("insert into test(name) value('clark')",tx)
	dialect.BatchExecute("insert into test(name) value('clark')",tx)
	dialect.BatchExecute("insert into test(name) value('clark')",tx)
	dialect.BatchExecute("insert into test(name) value('clark')",tx)
	dialect.Commit(tx)

	l , err := dialect.Query("select * from test")
	should.NoError(err)
	should.NotEqual(0,l.Len())

}

func Test_ConcurrentRequest(t *testing.T){
	dialect := Dialect{
		maxConns:40,
		maxIdles:30,
	}
	dialect.Create(driverSource)
	ResultTest(t,dialect)
}

func ResultTest(t *testing.T,dialect Dialect){
	var i = 0
	for{
		go func(){
			i++
			println(i)
			dialect.Query("select * from test")
		}()
	}
}

func Test_QueryResult(t *testing.T){
	dialect := Dialect{
		maxConns:100,
		maxIdles:50,
	}
	dialect.Create(driverSource)
	l , _ := dialect.Query("select * from test")
	for e := l.Front() ; e != nil ;e = e.Next() {
		val := e.Value.(map[string]interface{})
		fmt.Printf("key = %v , value = %v \n" ,val["id"], val["name"])
	}

	l , _ = dialect.Query("select count(*) as counts from test")

	m := l.Front().Value.(map[string]interface{})
	fmt.Printf("%v\n",m)

	// toInt
	i , err := dialect.ToInt("select count(*) as count from test")
	if err != nil {
		log.Error(err.Error())
	}
	log.Infof("Toint =%d \n",i)

	type s struct {
		Name string
		Id int
	}

	// toStruct
	ss := s{}
	r , err := dialect.ToStruct("select * from test",ss)
	if err != nil {
		log.Error(err.Error())
	}
	log.Infof("ToStruct , Name = %s , value = %d \n",r[0].(*s).Name ,r[0].(*s).Id)

	// toString
	r1 , err := dialect.ToString("select name from test limit 1")
	if err != nil {
		log.Error(err.Error())
	}
	log.Info(r1)

}


func init(){
	log.InitLog(1,0)
}