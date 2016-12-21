package main

import "fmt"

const n = 5

func left(a, b int) bool {
	return a+1 == b || a == n && b == 1
}
func right(a, b int) bool {
	return left(b, a)
}
func near(a, b int) bool {
	return left(a, b) || right(a, b)
}
func main() {
	zhao := 1
	for qian := 2; qian <= n; qian++ {
		for sun := 2; sun <= n; sun++ {
			if sun == qian {
				continue
			}
			for li := 2; li <= n; li++ {
				if li == qian || sun == qian {
					continue
				}
				zhou := n*(n+1)/2 - zhao - qian - sun - li
				if zhou < 0 || zhou > n || zhou == zhao || zhou == qian || zhou == sun || zhou == li {
					continue
				}
				zp, zq := near(zhao, qian), left(sun, qian) || left(li, qian)
				lp, lq := left(qian, sun), near(li, sun)
				if zp || zq || lp || lq {
					continue
				}
				fmt.Println(zhao, qian, sun, li, zhou)
			}
		}
	}
}
