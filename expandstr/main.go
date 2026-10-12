package main

import (
	"fmt"
	"os"
)

func main() {
	myArg := os.Args[1]
	newStr := ""
	if len(os.Args) > 2 {
		fmt.Print()
		return
	} else if myArg == "" {
		fmt.Println()
		return
	}

	for index, char := range myArg {
		if (char >= 'a' && char <= 'z') || char == '\'' {
			newStr += string(char)
		} else if index < len(myArg)-1 && (char == ' ' && index != 0) && myArg[index+1] >= 'a' && myArg[index+1] <= 'z' {
			newStr += string(char)
			newStr += " "

		}
	}
	fmt.Print(newStr)
}
