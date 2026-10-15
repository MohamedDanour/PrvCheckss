package main

import (
	"fmt"
	"os"
)
 
func IsPrime(input int) bool {
	if input == 1 {
		return false
	} else if input == 2 {
		return true
	} else if input <= 0 {
		return false
	}
	for i := 2; i*i <= input; i++ {
		if input%i == 0 {
			return false
		}
	}
	return true
}

func Sum_PreviousPrime(result int) int {
	Sum := 0
	for i := result; i >= 2; i-- {
		if IsPrime(i) {
			Sum += i
		} else {
			continue
		}
	}
	return Sum
}

func Atoi(k string) int {
	result := 0
	for i := 0; i < len(k); i++ {
		result = result*10 + int(k[i]) - '0'
	}
	return result
}

func main() {
	fmt.Print(Sum_PreviousPrime(Atoi(os.Args[1])))
}
