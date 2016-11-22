package main

import (
  "fmt"
  "sort"
)

const n = 100

var h = make([]int, n+1)

func fill() {
  for i := 0; i < len(h); i++ {
    h[i] = i * i * i * i * i
  }
}
func main() {
  fill()
  var a, b, c, d, e, f int
  for a = 1; a <= n; a++ {
    for b = a; b <= n; b++ {
      for c = b; c <= n; c++ {
        for d = c; d <= n; d++ {
          for e = d; e <= n; e++ {
            f = sort.SearchInts(h, h[a]+h[b]+h[c]+h[d]+h[e])
            fmt.Println("f=", f, "key=", h[a]+h[b]+h[c]+h[d]+h[e])
            if f > d && f <= n && h[f] == h[a]+h[b]+h[c]+h[d]+h[e] {
              fmt.Println(a, b, c, d, e, f)

            }
          }
        }
      }
    }
  }
}
