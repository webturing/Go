package main

import "fmt"

func Sqrt(x float64) float64 {
	z := float64(x)
	for i := 0; i < 10; i++ {
		z = z/2 + 1/z
	}
	return z
}
func main() {
	fmt.Printf("%g\n", Sqrt(2))
}
