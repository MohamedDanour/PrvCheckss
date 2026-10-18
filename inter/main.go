package inter
package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("")
		return
	}
	FirstARG := os.Args[1]
	SecondARG := os.Args[2]
	newStr := ""
	for _, char := range FirstARG {
		for i := 0; i <= len(SecondARG)-1; i++ {
			if char == rune(SecondARG[i]) {
				already := false
				for x := 0; x < len(newStr); x++ {
					if char == rune(newStr[x]) {
						already = true
					}
				}
				if !already {
					newStr += string(char)
				}
			}
		}
	}
	{
		for _, char2 := range newStr {
			z01.PrintRune(char2)
		}
	}

	z01.PrintRune('\n')
}
