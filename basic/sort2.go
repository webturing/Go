package main

import (
	"fmt"
	"sort"
)

func main() {
	intList := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	fmt.Println(intList)
	sort.Sort(sort.Reverse(sort.IntSlice(intList)))
	fmt.Println(intList)
}
