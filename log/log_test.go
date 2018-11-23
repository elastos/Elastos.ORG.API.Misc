package log

import (
	"testing"
)

func Test_Log(t *testing.T){

	InitLog(0,3)
	for {
		Info("helloworld")
	}

}

func Test_logLever(t *testing.T){
	InitLog(1,3)

	Debug("debug")
	Info( "info")
	Warn("warn")
	Error("error")

	Debugf("%s","debug")
	Infof("%s","info")
	Warnf("%s","warn")
	Errorf("%s","error")

}
