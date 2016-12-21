package main

import (
	"fmt"
	"strconv"
)

func main() {
	for n := 16; ; n += 10 {
		s := strconv.Itoa(n / 10)
		m, _ := strconv.Atoi("6" + s)
		if m == 4*n {
			fmt.Println(m, n)
			break
		}
	}
}
