package main

import (
	"fmt"
)

func RepeatAlpha(s string) string {

	newStr := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			count := char - 'a' + 1
			for i := 0; i <= int(count)-1; i++ {
				newStr += string(char)
			}
		} else if char >= 'A' && char <= 'Z' {
			count := char - 'A' + 1
			for i := 0; i <= int(count)-1; i++ {
				newStr += string(char)
			}
		} else {
			newStr += string(char)

		}

	}
	return newStr
}

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}
