package main

import (
	"github.com/elastos/Elastos.ORG.API.Misc/chain"
	"github.com/elastos/Elastos.ORG.API.Misc/http"
	"github.com/elastos/Elastos.ORG.API.Misc/log"
)

func main() {
	go chain.Sync()
	http.StartServer()
}

func init() {
	log.InitLog(0, 50)
}
