package tools

import "testing"

func Test_array_Contains(t *testing.T) {

	arr := []string{
		"test1",
		"test2",
	}

	println(Contains(arr, "asdf"))
	println(Contains(arr, 111))
	println(Contains(arr, "test1"))

	arr = nil

	println(Contains(arr, "asdf"))
	println(Contains(arr, 111))
	println(Contains(arr, "test1"))

}
