//a rough simple method to find cube root manually by exhaustive enumeration

package main

import (
  "math"
  "fmt"
  )

func Cuberoot(x float64) float64 {
  var count, ans, pos float64
  pos = math.Abs(x)
  for ; pos - ans*ans*ans >= 0 ; {
    ans += 1
    count += 1
  }
  if pos != x {
    ans = -ans
  }
  fmt.Println(count)
  return ans
}

//This func is so freaking close to real answer, like 7 digits after decimal!
//and its efficiency increases with bigger numbers!
func Cube_bisection(x float64) float64 {
  var count, temp, high, low, ans, pos float64

  pos = math.Abs(x)
  high = pos
  low = 0
  for {
    count += 1
    if high>low {
      ans = (high + low)/2
    }
    if ans == temp {
      if pos != x {
        ans = -ans
      }
      fmt.Println(count)
      return ans
    }
    temp = ans
    if (math.Pow(ans, 3) > pos) {
      high = ans
    } else {
      low = ans
    }
    //fmt.Println(ans)
  }
}

func main() {
  fmt.Println(Cuberoot(128975634354))
  fmt.Println(Cube_bisection(128975634354))
}
