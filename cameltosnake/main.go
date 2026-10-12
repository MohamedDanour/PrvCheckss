package main

import (
	"fmt"
)

func CamelToSnakeCase(s string) string {

	newString := ""
	if s == "" {
		return s
	}
	if len(s)-1 >= 'A' && len(s)-1 <= 'Z' {
		return s
	}
	if (s[0] >= 'A' && s[0] <= 'Z') && (s[1] >= 'A' && s[1] <= 'Z') {
		return s
	}

	if s[0] >= 'a' && s[0] <= 'z' || s[0] >= 'A' && s[0] <= 'Z' {
		for i := 0; i < len(s); i++ {
			if i > 0 && (s[i] >= 'A' && s[i] <= 'Z') && (s[i-1] >= 'a' && s[i-1] <= 'z') {
				newString += "_"
			}
			newString += string(s[i])
		}
	}
	return newString
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
