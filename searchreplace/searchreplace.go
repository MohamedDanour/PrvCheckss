package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Print("")
		return
	}
	FirstInput := os.Args[1]
	SecondInput := os.Args[2]
	ThirdInput := os.Args[3]
	newStr := ""
	for _, char := range FirstInput {
		if string(char) == SecondInput {
			newStr += ThirdInput
			continue
		}
		newStr += string(char)
	}
	fmt.Print(newStr)
}
