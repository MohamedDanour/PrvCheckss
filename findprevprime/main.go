package main

import "fmt"

func IsPrime(n int) bool {
	// 5 -  5 <= 10
	for i := 2; n <= i*i; i++ {
		if n%i != 0 {
			return false
		}
	}
	return true
}

func FindPrevPrime(nb int) int {
	prevprime := nb // 5

	if IsPrime(prevprime) { // true
		AnotherVar := prevprime - 1 // 4
		if IsPrime(AnotherVar) {
			return AnotherVar
		} else {
			return prevprime
		}
	} else {
		AnotherVar2 := prevprime - 1
		if !IsPrime(AnotherVar2) {
			return AnotherVar2
		} else {
			return prevprime
		}
	}

}

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))

}
