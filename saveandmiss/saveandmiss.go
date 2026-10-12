package main

import (
	"fmt"
)

func SaveAndMiss(arg string, num int) string {
	if num <= 0 {
		return arg
	}

	saved := ""
	Dontprint := num
	for i := 0; i <= len(arg)-1; i++ {
		if Dontprint == i {
			i += 3
			Dontprint += num * 2
		}
		saved += string(arg[i])
	}
	return saved
}
func main() {
	fmt.Println(SaveAndMiss("123456789", 3))
	fmt.Println(SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(SaveAndMiss("", 3))
	fmt.Println(SaveAndMiss("hello you all ! ", 0))
	fmt.Println(SaveAndMiss("what is your name?", 0))
	fmt.Println(SaveAndMiss("go Exercise Save and Miss", -5))
}

// abchinopqrstu vwyz
// abcghimnostuz
