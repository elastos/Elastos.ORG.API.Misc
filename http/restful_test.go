package http

import (
	"github.com/elastos/Elastos.ORG.API.Misc/log"
	"testing"
)

func Test_Start(t *testing.T){
	log.Debug("start server")
	StartServer()
}

func init(){
	log.InitLog(0,0)
}