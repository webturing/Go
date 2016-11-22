package main

import "fmt"

func h(x int) int {
  return x * x * x * x * x
}
func main() {
  var a, b, c, d, e, f int
  const n = 100
  for a = 1; a <= n; a++ {
    for b = a; b <= n; b++ {
      for c = b; c <= n; c++ {
        for d = c; d <= n; d++ {
          for e = d; e <= n; e++ {
            for f = d + 1; f <= n; f++ {
              if h(a)+h(b)+h(c)+h(d)+h(e) == h(f) {
                fmt.Println(a, b, c, d, e, f)
              }
            }
          }
        }
      }
    }
  }
}
