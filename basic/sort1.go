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
}
