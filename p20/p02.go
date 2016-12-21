package main

import "fmt"

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}
func lcm(a, b int) int {
	return a * b / gcd(a, b)

}
func sumr(a, b, r int) int {
	a, b = (a+r-1)/r*r, b/r*r
	n := (b-a)/r + 1
	return (a + b) * n / 2
}
func main() {
	fmt.Println(sumr(1, 1000, lcm(3, 5)))
}
