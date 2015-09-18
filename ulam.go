package main

import (
	"flag"
	"fmt"
  "math"
)

func main() {
	// body
  squire := flag.Int("d", 3, "3x3 squire")
  flag.Parse()
  fmt.Println( "squire", *squire )
  center := math.Ceil(float64(*squire)/2.0)
  fmt.Println( "starting cell", center )
  
  b := make([][]int, *squire)

  for i := 0; i < *squire; i++ {
    b[i] = make([]int, *squire)
    for j := 0; j < *squire; j++ {
      b[i][j] = 9
    }
  }
  
  for i := 0; i < *squire; i++ {
    fmt.Println( b[i] )
  }
  
}
