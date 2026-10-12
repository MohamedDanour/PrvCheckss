package main

import (
	"fmt"
)

func FirstWord(s string) string {
	newStr := ""
	if len(s) <= 0 {
		return "\n"
	}
	for index, char := range s {
		if char != ' ' {
			newStr += string(s[index])
		} else {
			break
		}
	}
	return newStr + "\n"
}

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}
