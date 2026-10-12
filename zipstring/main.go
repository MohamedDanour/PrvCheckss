package main

import (
	"fmt"
)

// check each letter
// as long as it not equal to the other letter count++
// save this number with the letter
// reset the counter

func ZipString(s string) string {
	var count int
	newString := ""
	if s == "" {
		return ""
	}
	if len(s) <= 1 {
		return "1" + s
	}
	for i := 1; i <= len(s)-1; i++ {
		count = 1
		if (len(s)-1 >= i && s[i] != s[i-1]) || (len(s)-1 >= i && s[i] != s[i-1] && s[i] != s[1+1]) {
			newString += "1" + string(s[i-1])
		}
		if s[i] == s[i-1] {
			K := i
			for K < len(s) && s[K] == s[K-1] {
				count++
				K++
				i++
			}
			newString += string(rune(count+'0')) + string(s[i-1])
		}
	}
	if s[len(s)-1] == s[len(s)-2] {
		return newString
	} else {
		return newString + "1" + string(s[len(s)-1])
	}
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
	//1Y1o3u1a2d1r1a1s1z
	//1Y1o3u1a2d1r1a1z
}

//1Y1o3u1n1g1F1e4l1a1s
//1T1h2e1 1q2u1i1c1k1 1b1r1o2w1n1 1f1o1x1 1j2u1m1p1s1 1o1v1e1r1 1t1h1e1 1l3a1z1y1 1d1o1g
//1H1e2l2o1 1T1h1e2r1e1!

//1Y1o3u1n1g1F1e4l1a1s
//1T1h2e1 1q2u1i1c1k1 1b1r1o2w1n1 1f1o1x1 1j2u1m1p1s1 1o1v1e1r1 1t1h1e1 1l3a1z1y1 1d1o1g
//1H1e2l2o1 1T1h1e2r1e1!
