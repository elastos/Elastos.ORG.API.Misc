package tools

import "testing"

func Test_bytesReverse(t *testing.T) {

	b := []byte{}

	ReverseBytes(b)

	b = []byte{
		1,
	}

	ReverseBytes(b)

	b = []byte{
		1,
		2,
	}

	ReverseBytes(b)

	b = []byte{
		1,
		2,
		3,
	}

	ReverseBytes(b)
}
