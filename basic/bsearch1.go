package main

import (
	"fmt"
	"sort"
)

func main() {
	intList := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	fmt.Println(intList)
	sort.Ints(intList)
	fmt.Println(intList)
	pos := sort.SearchInts(intList, 7)
	if pos < len(intList) && pos >= 0 {
		fmt.Println(pos)
	} else {
		fmt.Println("not found")
	}
}
