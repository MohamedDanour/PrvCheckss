package main

import "fmt"

func CountRepeats(s string) int {
	counter := 0
	for i := 0; i < len(s); i++ {
		found := false

		for j := 0; j < i; j++ {
			if s[i] == s[j] {
				found = true
				break
			}
		}
		if found {
			continue
		}
		for k := i + 1; k < len(s); k++ {
			if s[i] == s[k] {
				counter++
				break
			}
		}
	}
	return counter
}

func main() {
	fmt.Println(CountRepeats("hello"))
	fmt.Println(CountRepeats("abc"))
	fmt.Println(CountRepeats("aabbccd"))
	fmt.Println(CountRepeats("aaa"))

}
