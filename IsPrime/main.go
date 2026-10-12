package main

import "fmt"

func IsPrime(n int) bool {
	if n == 1 {
		return false
	} else if n == 2 {
		return true
	} else if n <= 0 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func NextPrime(n int) int {
	the_next_num := n
	if n >= 0 {
		the_next_num = n + 1
		for !IsPrime(the_next_num) { // false
			the_next_num++
		}
		return the_next_num
	} else {
		return 0
	}
}

func Sum_PreviousPrime(n int) int {
	Sum := 0
	for i := n; i >= 2; i-- {
		if IsPrime(i) {
			Sum += i
		} else {
			continue
		}
	}
	return Sum
}

func fprime(n int) int {

}

func main() {
	fmt.Println(Sum_PreviousPrime(79))

}
