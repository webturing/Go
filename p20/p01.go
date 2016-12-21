package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	start, end, target := 1, 100, 50
	for {
		cur := r.Intn(end-start+1) + start
		fmt.Print(cur, " ")
		if cur == target {
			break
		}
	}
}
