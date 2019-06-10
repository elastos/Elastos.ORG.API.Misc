package tools

import "testing"

func Test_getAddress(t *testing.T) {

	str := "02512bf24fe87a7eb6d831f80a243721e4da13735a4f1e5f2cc717b20f1974645c"

	addr, _ := GetAddress(str)
	println(addr)
}
