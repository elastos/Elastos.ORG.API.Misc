package chain

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func Test_Hex(t *testing.T) {
	//println(hexutil.EncodeUint64(10))
	//
	//b, _ := hexutil.DecodeUint64("0xa")
	//println(b)
	//
	//println(bytesTo32Int([]byte("12")))

	b := []byte{35}
	bin_buf := bytes.NewBuffer(b)
	var x byte
	binary.Read(bin_buf, binary.BigEndian, &x)
	println(int(x))

}

// bytes to int 32
func bytesTo32Int(b []byte) int {
	buf := bytes.NewBuffer(b)
	var tmp uint32
	binary.Read(buf, binary.BigEndian, &tmp)
	return int(tmp)
}
