package main

import (
	"fmt"
	solution "./solution"
)

func main()  {
	// Number one
	str := "okaaryanataokaokaaryanataoka"
	fmt.Println(solution.NumberOne(str,5))
	
	// Number two
	// slice := []int{1,1,2}
	slice := []int{0,0,1,1,1,2,2,3,3,4}
	fmt.Println(solution.RemoveDuplicate(slice))

	// Number Four
	binStr := "0000001011"
	binStr = "11111111111111111111111111111101"
	fmt.Println(solution.HammingWeight(binStr))
	
}