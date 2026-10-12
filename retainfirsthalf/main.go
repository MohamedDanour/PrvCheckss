package main

import (
	"fmt"
)

func RetainFirstHalf(str string) string {
	half := len(str) / 2
	newStr := ""
	if str == "" {
		return ""
	} else if len(str) == 1 {
		return str
	}
	for i := 0; i < half; i++ {
		newStr += string(str[i])
	}
	return string(newStr)
}

func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
}
