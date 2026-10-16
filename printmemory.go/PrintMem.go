package main

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	hex := "0123456789abcdef"
	for i, x := range arr {
		z01.PrintRune(rune(hex[x/16]))
		z01.PrintRune(rune(hex[x%16]))
		if i == len(arr)-1 || i%4 == 3 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}

	for _, y := range arr {
		if y >= 32 && y <= 132 {
			z01.PrintRune(rune(y))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}
func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
