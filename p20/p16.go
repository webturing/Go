package main

import (
	"fmt"
)

func main() {
	for n := 10; n <= 99; n++ {
		a, b := n/10, n%10
		if n%(a+b) == 0 {
			fmt.Println(n)
		}
	}
}
