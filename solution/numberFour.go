package solution

func HammingWeight(binStr string) (result int) {
	result = 0
	for _, val := range binStr {
		if string(val) == "1" {
			result += 1
		}
	}
	return
}
