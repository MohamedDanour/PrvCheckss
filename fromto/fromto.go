package main

import (
	"fmt"
)

func FromTo(from int, to int) string {
	newStr := ""

	if from == to {
		return string(rune(to/10+'0')) + string(rune(to%10+'0')) + "\n"
	} else if from > 99 || to > 99 {
		return "invalid"
	}

	for from <= to {
		newStr += "0" + string(rune(from+'0')) + "," + " "
		from++
		if from == to {
			newStr += string(rune(to/10+'0')) + "0"
			from--
			break
		}

	}

	for from >= to {
		newStr += "0" + string(rune(to+'0')) + "," + " "
		to++
		if from == to {
			newStr += string(rune(from/10+'0')) + "0"
			break
		}
	}

	return newStr + "\n"
}

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(100, 10))

}
