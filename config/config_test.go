package config

import (
	"github.com/elastos/Elastos.ORG.API.Misc/log"
	"testing"
)

func Test_Config(t *testing.T){
	println(1)
	log.Info(Conf.ServerPort,Conf.Node,Conf.DbDriverName,Conf.DbDriverSource)
}

func init(){
	log.InitLog(1,0)
}