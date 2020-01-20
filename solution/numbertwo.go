package solution

func RemoveDuplicate(slice []int) []int {
	res := []int{}
	prev := 0
	lenSlice := len(slice)
	if lenSlice != 0 {
		prev = slice[0]
		res = append(res, prev)
	}
	for x := 1; x < lenSlice; x++ {
		if slice[x] != prev {
			res = append(res, slice[x])
			prev = slice[x]
		}
	}
	return res
}
