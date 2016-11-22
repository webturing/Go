package main

import "fmt"

func h(x int) int {
  return x * x * x * x * x
}
func main() {
  var a, b, c, d, e, f int
  const n = 100
  for a = 1; a <= n; a++ {
    for b = 1; b <= a; b++ {
      for c = 1; c <= b; c++ {
        for d = 1; d <= c; d++ {
          for e = 1; e <= d; e++ {
            for f = a + 1; f <= n; f++ {
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
