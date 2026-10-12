package main

import (
	"fmt"
)

func LastWord(s string) string {

	i := len(s) - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}

	Thelastchar := i

	for i >= 0 && s[i] != ' ' {
		i--
	}

	start := i + 1
	newStr := ""

	for start <= Thelastchar {
		newStr += string(s[start])
		start++
	}
	return newStr + "\n"
}

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}
