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
	Eth        ETH
	Cmc        Cmc
	EnableCors bool
}

type Cmc struct {
	ApiKey    []string
	Inteval   string
	NumOfCoin int
	Enable    bool
}

type DB struct {
	DbDriverName   string
	DbDriverSource string
}

type ELA struct {
	Restful         string
	Jsonrpc         string
	JsonrpcUser     string
	JsonrpcPassword string
	Enable          bool
}

type BTC struct {
	Host       string
	Rpcuser    string
	Rpcpasswd  string
	MinConfirm int
	Net        string
	Enable     bool
}

type ETH struct {
	Endpoint    string
	Enable      bool
	StartHeight int64
	BatchSize   int
	Retry       int
	SideChain   bool
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
