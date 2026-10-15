package main

import "fmt"

func IsPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func FindPrevPrime(nb int) int {

	if IsPrime(nb) {
		return nb
	} else {
		for !IsPrime(nb) {
			nb--
		}
	}
	return nb
}
func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(15))
}
