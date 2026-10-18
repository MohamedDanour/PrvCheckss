package main

import (
	"github.com/01-edu/z01"
)

func Chunk(slice []int, size int) {
	var myArr [][]int

	if size == 0 {
		z01.PrintRune('\n')
		return
	}

	if len(slice) == 0 {
		z01.PrintRune('[')
		z01.PrintRune(']')
		z01.PrintRune('\n')
		return
	}

	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		myArr = append(myArr, slice[i:end])
	}

	for index, Chunk := range myArr {
		z01.PrintRune('[')
		if index == 0 {
			z01.PrintRune('[')
		}
		for i, num := range Chunk {

			for _, char := range itoa(num) {
				z01.PrintRune(char)
			}
			if i == len(Chunk)-1 {
				continue
			} else {
				z01.PrintRune(' ')

			}
		}
		if index != len(myArr)-1 {
			z01.PrintRune(']')
			z01.PrintRune(' ')
		} else {
			z01.PrintRune(']')
		}
	}

	z01.PrintRune(']')
	z01.PrintRune('\n')
}
func itoa(n int) string {
	myStr := ""
	if n == 0 {
		converted := string(rune(n + '0'))
		return converted
	}

	for n != 0 {
		newInt := n % 10
		n = n / 10
		converted := string(rune(newInt + '0'))
		myStr = converted + myStr
	}
	return myStr
}

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
// to fit the bloody z01 rules 