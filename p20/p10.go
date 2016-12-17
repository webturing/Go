package main

import (
	"fmt"
)

func bool2int(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}
func main() {
	for a := 1; a <= 3; a++ {
		for b := 1; b <= 3; b++ {
			for c := 1; c <= 3; c++ {
				if a != b && b != c && c != a {
					ap, aq := a == 1, b == 2 && c == 2
					bp, bq := b <= 2, c == 3
					cp, cq := a != 2, b != 1
					if a+bool2int(ap)+bool2int(aq) == 3 && b+bool2int(bp)+bool2int(bq) == 3 && c+bool2int(cp)+bool2int(cq) == 3 {
						fmt.Println(a, b, c)
					}
				}
			}
		}
	}
}
