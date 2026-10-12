package main

import (
	"fmt"
	"math"
)

func PasswordEntropy(pw string) string {
	if pw == "" {
		return "invalid input"
	}
	var pool int
	hasLower, hasUpper, hasDigit, hasOther := false, false, false, false
	for _, char := range pw {
		if char >= 'a' && char <= 'z' {
			hasLower = true
		} else if char >= 'A' && char <= 'Z' {
			hasUpper = true
		} else if char >= '1' && char <= '9' {
			hasDigit = true
		} else {
			hasOther = true
		}

	}

	if hasLower {
		pool += 26
	}
	if hasUpper {
		pool += 26
	}
	if hasDigit {
		pool += 10
	}
	if hasOther {
		pool += 32
	}
	Entropy := len(pw) * int(math.Log2(float64(pool)))
	if Entropy <= 40 {
		return "weak" + "\n"
	} else if Entropy <= 60 {
		return "medium" + "\n"
	} else {
		return "strong" + "\n"
	}
}

func main() {
	fmt.Print(PasswordEntropy("abc"))
	fmt.Print(PasswordEntropy("Password1"))
	fmt.Print(PasswordEntropy("C0rrect-Horse-B4ttery"))
	fmt.Print(PasswordEntropy(""))
}
