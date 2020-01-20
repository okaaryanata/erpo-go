package solution

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func NumberOne(str string, k int) string {
	maxVal := 2000
	if len(str) <= maxVal {
		tmp := Reverse(str[:k])
		return tmp
	} else {
		counter := len(str)
		res := ""
		y := 0
		for x := 1; x < counter+1; x++ {
			if counter > maxVal*x {
				tmp := str[(maxVal * (x - 1)):(maxVal * x)]
				res += Reverse(tmp[:k]) + tmp[k:]
				y = x
			}
		}
		res += Reverse(str[maxVal*y:])
		return res
	}
}
