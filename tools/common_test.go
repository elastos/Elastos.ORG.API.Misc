package tools

import "testing"

func Test_getAddress(t *testing.T) {

	str := "038f14acf248383e33aa8b098a60ed20184a66fdcef018d0eeb0dbd816a4ca0ddd"

	addr, _ := GetAddress(str)
	println(addr)
}
