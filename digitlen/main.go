package main

import (
	"fmt"
)

func DigitLen(n, base int) int {
	var result int
	if base < 0 {
		return -1
	}
	for n != 0 {
		n = n / base
		result++
	}
	return result
}

func main() {
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
}
