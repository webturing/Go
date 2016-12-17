package main

import "fmt"

func prime(n int) bool {
	if n == 2 {
		return true
	}
	if n < 2 || n%2 == 0 {
		return false
	}
	for c := 3; c*c <= n; c += 2 {
		if n%c == 0 {
			return false
		}
	}
	return true
}
func main() {
	for n := 100; n <= 999; n++ {
		head, mid, tail := n/100, (n%100)/10, n%10
		if head != mid && mid != tail && tail != head && mid > head+tail && !prime(mid+head) {
			fmt.Print(head*100+mid*10+tail, " ")
		}
	}
}
