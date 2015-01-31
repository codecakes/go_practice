//successive approximation
package main

import (
  "fmt"
  //"math"
  )

  func Sqrt(x float64) float64 {
    z := float64(1)
    temp := z
    for {
      temp = z -  ((z*z) - x)/(2*z)
      if (temp - z < 0.0000000001) {
        return temp
      }
      z = temp
    }
  }

  func main() {
    fmt.Println(Sqrt(2))
  }
