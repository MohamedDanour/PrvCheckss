package main

import (
	"fmt"
)

func HashCode(dec string) string {
	newSTR := ""
	for i := 0; i <= len(dec)-1; i++ {

		newChar := int(int(dec[i])+len(dec)) % 127
		if newChar == 127 && newChar <= 32 {
			newChar += 33
		}
		newSTR += string(newChar)
	}

	return newSTR
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}

// $ go run .
