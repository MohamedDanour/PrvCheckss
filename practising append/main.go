package main

import "fmt"

func Evens(n int) []int {

	newArr := []int{}

	for i := 2; i <= n; i++ {

		if i%2 == 0 {

			newArr = append(newArr, i)

		} else if n%i != 0 {

			continue

		} 

	}

	return newArr

}

func main() {
	fmt.Println(Evens(10))
	fmt.Println(Evens(1))
	fmt.Println(Evens(6))
}
