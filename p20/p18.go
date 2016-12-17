package main

import (
	"fmt"
)

func one(n int) int {
	ret := 0
	for n != 0 {
		if n%10 == 1 {
			ret++
		}
		n /= 10
	}
	return ret
}
func main() {
	var tot, p = 0, 0
	for p = 1; tot < 202; p++ {
		tot += one(p)
	}
	if tot == 202 {
		fmt.Println(p)
	} else {
		fmt.Println("Impossible")
	}
}
