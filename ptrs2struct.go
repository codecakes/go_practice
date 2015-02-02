//just playing out with ptrs to structs

package main

import "fmt"

type Vertex struct {
  X int
  Y int
}

func main() {
  v := Vertex{1, 2}
  p := &v
  //&{1 2}
  fmt.Println(p)
  //{1 2}
  fmt.Println(*p)
  p.X = 1e9
  fmt.Println(v)
}
