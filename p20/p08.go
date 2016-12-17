package main

import "fmt"

func imply(p, q bool) bool {
	return !p || q
}
func main() {
	for k := 0; k < 32; k++ {
		a, b, c, d, e := k/16, (k%16)/8, (k%8)/4, (k%4)/2, (k % 2)
		if a+b+c+d+e == 3 && a+c <= 2 && b+c > 0 && imply(c == 1, d+e == 1) && imply(b == 1, d+e < 2) {
			fmt.Println(a, b, c, d, e)
		}
	}

}
