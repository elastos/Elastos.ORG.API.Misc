package tools

func Contains(src []string,target interface{}) bool{

	for _ , v := range src {
		if v == target {
			return true
		}
	}

	return false
}