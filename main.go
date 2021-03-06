package main

import (
	"github.com/elastos/Elastos.ORG.API.Misc/chain"
	"github.com/elastos/Elastos.ORG.API.Misc/config"
	"github.com/elastos/Elastos.ORG.API.Misc/http"
)

func main() {
	if config.Conf.Ela.Enable {
		go chain.Sync()
	}
	if config.Conf.Eth.Enable {
		go chain.SyncEth()
	}
	http.StartServer()
}
