package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 100
	h := make([]int, n+1)
	for i := 1; i <= n; i++ {
		h[i] = i * i * i * i * i
	}
	k := 100001
	m := sort.SearchInts(h, k)
	if h[m] == k {
		fmt.Println("find")
	} else {
		fmt.Println("not found")
	}

}
