package main

import "fmt"

const n = 100

var h = make([]int, n+1)

func fill() {
  for i := 0; i < len(h); i++ {
    h[i] = i * i * i * i * i
  }
}
func main() {
  fill()
  fmt.Println(h)
  var a, b, c, d, e, f int
  for a = 1; a <= n; a++ {
    for b = 1; b <= a; b++ {
      for c = 1; c <= b; c++ {
        for d = 1; d <= c; d++ {
          for e = 1; e <= d; e++ {
            for f = a + 1; f <= n; f++ {
              if h[a]+h[b]+h[c]+h[d]+h[e] == h[f] {
                fmt.Println(a, b, c, d, e, f)
              }
            }
          }
        }
      }
    }
  }
}
