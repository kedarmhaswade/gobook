package main

import (
  "os"
  "fmt"
)

func main() {
  fmt.Println(os.Args[1:])
  /*
  s, sep := "", ""
  for _, arg := range os.Args[1:] {
    s += sep + arg
    sep = " "
  }
  fmt.Println(s)
  */
}
