package tools

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(fullPath string) (b []byte) {

	_, err := os.Stat(fullPath)
	if err != nil {
		fmt.Printf("Can not find %s ", fullPath)
		return nil
	}

	b, err = ioutil.ReadFile(fullPath)
	if err != nil {
		fmt.Printf("Read file err %s ", err.Error())
	}
	return
}
