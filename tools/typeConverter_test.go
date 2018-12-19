package tools

import (
	"fmt"
	"testing"
)

func Test_typeConverter(t *testing.T) {

	var tt struct {
		Name string
		Age  int64
	}

	m := map[string]interface{}{
		"Name": "clark",
		"Age":  10,
	}

	Map2Struct(m, &tt)

	fmt.Printf("%v", tt)

}
