package main

import (
	"fmt"
	"os"
)

func main() {

	FirstInput := os.Args[1:]

	for _, arg := range FirstInput {
		newStr := ""

		for _, char := range arg {
			if char >= 'A' && char <= 'Z' {
				newStr += string(char + 32)
			} else if char >= 'a' && char <= 'z' || char == ' ' || char == '\'' || char >= '0' && char <= '9' {
				newStr += string(char)
			}
		}
		AnoterNewStr := ""

		for index, char := range newStr {

			if index+1 < len(newStr) && char >= 'a' && char <= 'z' && newStr[index+1] == ' ' {
				AnoterNewStr += string(char - 32)
			} else if index == len(newStr)-1 && char != ' ' {
				AnoterNewStr += string(char - 32)
			} else {
				AnoterNewStr += string(char)
			}

		}
		fmt.Println(AnoterNewStr)

	}

}

// seconD tesT iS A littlE easieR$
// bewarE it'S noT harD wheN $
//  gO A dernieR 0123456789 foR thE roaD E$

//  seconD tesT iS A littlE easieR$
// bewarE it'S noT harD wheN $
//  gO A dernieR 0123456789 foR thE roaD E$
