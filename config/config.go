package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type config struct {
	Node string
	ServerPort string
	DbDriverName string
	DbDriverSource string
}

var Conf *config

func init(){
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