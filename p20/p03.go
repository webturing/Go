package main

import "fmt"

var n int = 9

func main() {
	for i := 1; i <= n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		for j := i - 1; j >= 1; j-- {
			fmt.Print(j)p0
		}
		fmt.Println()
	}
	for i := n - 1; i >= 1; i-- {
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		for j := i - 1; j >= 1; j-- {
			fmt.Print(j)
		}
		fmt.Println()
	}
}
