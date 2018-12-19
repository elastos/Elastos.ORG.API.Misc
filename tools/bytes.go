package tools

//ReverseBytes reverse the order of the given bytes
func ReverseBytes(b []byte) []byte {
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
	return b
}
