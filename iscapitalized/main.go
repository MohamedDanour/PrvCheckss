package iscapitalized
package main

import (
	"fmt"
)

func IsCapitalized(s string) bool {
	if s == "" {
		return false
	}
	x := 0
	if s[x] >= 'a' && s[x] <= 'z' {
		return false
	}

	for index, char := range s {
		if char == ' ' { 
		if index < len(s)-1 && s[index+1] >= 'a' && s[index+1] <= 'z'  {
				return false 
		}
		}
	}
	return true

}
func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}
