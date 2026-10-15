package main

import (
	"fmt"
)

func PrintMemory(arr [10]byte) {

	for i, b := range arr {
		fmt.Print("%02x", b)
		if i == len(arr)-1 || i%4 == 3 {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}
	}
	for _, b := range arr {
		if b >= 32 && b <= 126 {
			fmt.Print(string(rune(b)))
		}
	}
	fmt.Print(".")
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
