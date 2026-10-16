package main

import (
	"github.com/01-edu/z01"
)

func main() {
	var First_loop int
	var Second_loop int
	var Third_loop int
	for First_loop = 9; First_loop >= 2; First_loop-- {
		for Second_loop = 8; Second_loop >= 1; Second_loop-- {
			for Third_loop = 7; Third_loop >= 0; Third_loop-- {
				if (First_loop > Second_loop) && (Second_loop > Third_loop) {
					z01.PrintRune(rune(First_loop + '0'))
					z01.PrintRune(rune(Second_loop + '0'))
					z01.PrintRune(rune(Third_loop + '0'))
					if !(First_loop == 2 && Second_loop == 1 && Third_loop == 0) {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}

			}
		}
	}
}
