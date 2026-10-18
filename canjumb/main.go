package canjumb
package main

import "fmt"

func CanJump(n []uint) bool {
	
	pos := 0
	for pos < len(n)-1 {
		steps := int(n[pos])
		if steps == 0 {
			return false
		}
		pos += steps
		if pos >= len(n) {
			return false
		}
	}
	return pos == len(n)-1
}

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))
}
