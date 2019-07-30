package chain

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"testing"
)

func Test_Hex(t *testing.T) {
	println(hexutil.EncodeUint64(10))

	b, _ := hexutil.DecodeUint64("0xa")
	println(b)

}
