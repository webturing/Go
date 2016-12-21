package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "12345678"
	n := len(s)
	w := (n + 2) / 3
	head, _ := strconv.Atoi(s[0 : n-2*w])
	mid, _ := strconv.Atoi(s[n-2*w : n-w])
	tail, _ := strconv.Atoi(s[n-w:])
	fmt.Println(head, mid, tail, head+tail-mid)
}
