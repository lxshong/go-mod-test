package go_mod_test

func Sum(params... int) int {
	var sum = 0
	for _, param := range params {
		sum += param
	}
	return sum
}
