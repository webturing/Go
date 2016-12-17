package main

import "fmt"

func search(total, money, fx, fy, fz float64) {
	for x := 0; x <= int(total); x++ {
		for y := 0; y <= int(total)-x; y++ {
			z := total - float64(x+y)
			if float64(x)*fx+float64(y)*fy+float64(z)*fz == money {
				fmt.Println(x, " ", y, " ", z)
			}
		}
	}
}
func main() {
	search(100, 100, 5, 3, 1./2)
	search(100, 100, 2.0, 4.0, 2.0/9.0)
	search(50, 30, 3, 2, 1)
}
