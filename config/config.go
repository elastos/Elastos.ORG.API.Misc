package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type config struct {
	ServerPort string
	Db         DB
	Ela        ELA
	Btc        BTC
	Cmc        Cmc
	VisitKey   string
}

type Cmc struct {
	ApiKey    []string
	Inteval   string
	NumOfCoin int
}

type DB struct {
	DbDriverName   string
	DbDriverSource string
}

type ELA struct {
	Host string
}

type BTC struct {
	Host       string
	Rpcuser    string
	Rpcpasswd  string
	MinConfirm int
	Net        string
}

var Conf *config

func init() {
	Conf = new(config)
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&Conf)
	if err != nil {
		fmt.Println("Error init Config :", err.Error())
		os.Exit(-1)
	}
}
