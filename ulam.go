package main

import (
	"flag"
	"fmt"
)

func main() {
	// body
  squire := flag.Int("d", 3, "3x3 squire")
  flag.Parse()
  fmt.Println( "squire", *squire )
  fmt.Println( "tail:", flag.Args() )

}
