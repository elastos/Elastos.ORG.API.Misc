package config

import (
	"github.com/elastos/Elastos.ORG.API.Misc/log"
	"testing"
)

func Test_Config(t *testing.T) {
	println(1)
	var c = Conf
	log.Info(c,Conf.ServerPort, Conf.Db.DbDriverSource, Conf.Btc.Host)
}

func init() {
	log.InitLog(1, 0)
}
