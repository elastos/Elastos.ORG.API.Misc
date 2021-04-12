package chain

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/elastos/Elastos.ORG.API.Misc/tools"
	"reflect"
	"testing"
)

func Test_SyncChain(t *testing.T) {
	Sync()
}

func Test_handleMemo(t *testing.T) {
	jsonStr := `{
      "msg": "7B22546167223A224449442050726F7065727479222C22566572223A22312E30222C22537461747573223A312C2250726F70657274696573223A5B7B224B6579223A224E616D65222C2256616C7565223A2261736466617364666166222C22537461747573223A317D5D7D",
      "pub": "02B536B5BC083883CF645ED60006AEB421575CA536C152366DF8F1085C7CCD7547",
      "sig": "C00D37AF1CD8C17CB2B2B82967E1C4F56EA2D4DEF74A5D345ACD06B7C247272BDAF6F6AE7053250DCA0C10C3AD212473E3C62DBBF3E8A78952DB19C9B5EBC267"
    }`
	hexStr := hex.EncodeToString([]byte(jsonStr))
	tx, _ := dba.Begin()
	err := handleMemo(hexStr, 10, "asdfafda", 190281, tx)
	if err != nil {
		dba.Rollback(tx)
	}
	dba.Commit(tx)

	a := 0

	if true {
		a := 1
		println(a)
	}

	println(a)

}

func Test_TypeConverter(t *testing.T) {

	header := Block_header{}
	result := make(map[string]interface{})
	jsonStr := `{"hash":"71b422e09dcd2f749d2adc0086735c210084cdb6b59bd4cd42e50455d024a662",
"confirmations":196504,"strippedsize":498,"size":498,"weight":1992,"height":1,"version":0,"minerinfo":"ELA"}`
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	tools.Map2Struct(result, &header)

	fmt.Println(header)
}

func Test_(t *testing.T) {
	var b float64 = 100000000
	a := b / ELA
	println(a, reflect.TypeOf(a).Name())
}
