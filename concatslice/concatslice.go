package main

import (
	"fmt"
)

func ConcatSlice(slice1, slice2 []int) []int {
	if len(slice1) == 0 {
		return slice2
	} else if len(slice2) == 0 {
		return slice1
	}

	res := []int{}

	for i := 0; i <= len(slice1)-1; i++ {
		res = append(res, slice1[i])

	}
	for x := 0; x <= len(slice2)-1; x++ {
		res = append(res, slice2[x])

	}
	return res
}
func main() {
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{}))
}
