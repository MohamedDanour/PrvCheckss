package main

import (
	"fmt"
)

func main() {

	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))

}

func FromTo(from int, to int) string {

	if from >= 99 || to >= 99 {
		return "invalid"
	}

	if from >= to {
		for i := to; i <= from; i++ {
			if i == 10 {
				return "10" + "\n"
			} else if from == to {
				fmt.Print(from)
				break
			}
			fmt.Print("0")
			fmt.Print(i)
			fmt.Print(",")
			fmt.Print(" ")
		}
	}

	if from <= to {
		for x := to; x >= from; x-- {

			if from == to {
				fmt.Print(from)
				break
			}
			if x == 10 {
				fmt.Print(x)
				fmt.Print(",")
				fmt.Print(" ")
			} else {
				fmt.Print("0")
				fmt.Print(x)
				if x != 1 {
					fmt.Print(",")
					fmt.Print(" ")

				}

			}

		}

	}

	return "\n"
}
