package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi(input string) int {
	result := 0
	for _, char := range input {
		result = result*10 + int(char-'0')
	}
	return result
}

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func sumprev(n int) int {
	sum := 0
	for n >= 2 {
		if isPrime(n) {
			sum += n
			n--
		}
		if !isPrime(n) {
			n--
		}
	}
	return sum
}

func printchar(n int) {
	if n < 0 {
		z01.PrintRune('0')
		z01.PrintRune('\n')

		return
	}

	if n == 0 {
		return
	}
	printchar(n / 10)
	z01.PrintRune(rune(n%10) + '0')
}

func main() {

	if len(os.Args) != 2 {
		z01.PrintRune('0')
		z01.PrintRune('\n')

		return
	}
	input := Atoi(os.Args[1])
	sum := sumprev(input)
	if sum == 0 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	printchar(sum)
	z01.PrintRune('\n')

}
