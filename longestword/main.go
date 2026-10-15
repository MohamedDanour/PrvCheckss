package main

import "fmt"

func LongestWord(s string) string {
	newArr := []string{}
	word := ""
	index := 0
	for len(s) > index {
		for index <= len(s)-1 && s[index] != ' ' {
			word += string(s[index])
			index++
		}

		if word == "" {
			index++
			continue
		}

		newArr = append(newArr, word)
		word = ""
		index++
	}

	Newstring := ""
	for x := 0; x <= len(newArr)-1; x++ {
		if len(Newstring) < len(newArr[x]) {
			Newstring = newArr[x]
		}
	}
	return Newstring
}

func main() {
	fmt.Printf("%q\n", LongestWord("I love programming in Go"))
	fmt.Printf("%q\n", LongestWord("  ab   cd  "))
	fmt.Printf("%q\n", LongestWord(""))
}
