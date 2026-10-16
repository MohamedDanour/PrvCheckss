package main

import (
	"fmt"
)

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}
	x := len(str1)
	z := len(str2)
	for i := 0; i <= len(str1)-1; i++ {
		if str1[i] == str2[i] {
			x--
			z--
		}
	}
	result := x + z
	return result
}

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
}
