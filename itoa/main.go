package main

import (
	"fmt"
)

func Itoa(n int) string {
	MyStr := ""
	if n == 0 {
		convert := string(rune(n + '0'))
		MyStr += convert
		return MyStr

	}

	z := n
	for z != 0 {
		if z < 0 {
			z = -z
		}
		newINT := z % 10
		z = z / 10
		converted := string(rune(newINT + '0'))
		MyStr = converted + MyStr
	}
	if n < 0 {
		MyStr = "-" + MyStr
	}
	return MyStr
}

func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}
