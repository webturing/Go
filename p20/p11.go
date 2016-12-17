package main

import "fmt"

const N = 64
const M = N / 4

func main() {
	for a := N / 2; a <= N; a++ {
		for b := 0; b <= N-a; b++ {
			for c := 0; c <= N-a-b; c++ {
				d := N - a - b - c
				A, B, C, D := a, b, c, d
				A, B, C, D = A-B-C-D, 2*B, 2*C, 2*D
				B, C, D, A = B-A-C-D, 2*C, 2*D, 2*A
				C, D, A, B = C-A-B-D, 2*D, 2*A, 2*B
				D, A, B, C = D-A-B-C, 2*A, 2*B, 2*C
				if A == M && B == M && C == M {
					fmt.Println(a, b, c, d)
				}
			}
		}
	}
}
