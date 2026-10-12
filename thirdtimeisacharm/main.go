package main

import (
	"fmt"
)

func ThirdTimeIsACharm(str string) string {
	newStr := ""

	for i := 2; i <= len(str)-1; i += 3 {
		newStr += string(str[i])
	}
	return newStr + "\n"
}

func main() {
	fmt.Print(ThirdTimeIsACharm("123456789"))
	fmt.Print(ThirdTimeIsACharm(""))
	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(ThirdTimeIsACharm("12"))
}
