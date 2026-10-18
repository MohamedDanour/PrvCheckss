package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("")
		return
	}
	FirstARG := os.Args[1]
	SecondARG := os.Args[2]

	i := 0
	foundd := false
	for _, char := range FirstARG {

		for ; i <= len(SecondARG)-1; i++ {
			foundd = false

			if char == rune(SecondARG[i]) {
				foundd = true
				i++
				break
			}
		}
	}

	if foundd {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}

}
