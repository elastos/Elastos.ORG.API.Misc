package tools

import (
	"github.com/elastos/Elastos.ORG.API.Misc/log"
	"io/ioutil"
	"os"
)

func ReadFile(fullPath string) (b []byte){

	_ , err := os.Stat(fullPath)
	if err != nil {
		log.Warnf("Can not find %s ",fullPath)
		return nil
	}

	b , err = ioutil.ReadFile(fullPath)
	if err != nil {
		log.Warnf("Read file err %s ",err.Error())
	}
	return
}