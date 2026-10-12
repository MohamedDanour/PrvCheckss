package main

import (
	"fmt"
)

func CountChar(str string, c rune) int {
	var counter int
	for _, char := range str {
		if char == c {
			counter++
		}
	}
	return counter
}

func main() {
	fmt.Println(CountChar("Hello World", 'l'))
	fmt.Println(CountChar("5  balloons", 5))
	fmt.Println(CountChar("   ", ' '))
	fmt.Println(CountChar("The 7 deadly sins", '7'))
}
