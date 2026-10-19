package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println()
		return
	}
	input := os.Args[1]
	newStr := ""
	for index, char := range input {
		if char >= 'a' && char <= 'z' || char == '\'' {
			newStr += string(char)
		}
		if char == ' ' {
			if len(newStr) > 0 && index < len(input)-1 && (input[index+1]) >= 'a' && (input[index+1] <= 'z') {
				newStr += " "
				newStr += " "
				newStr += " "
			}
		}
	}
	fmt.Println(newStr)
}
