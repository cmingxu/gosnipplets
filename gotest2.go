package main

import "fmt"

func adder(a, b int) int {
  return a + b
}

func main() {
  var a int = 1
  var b int = 2

  fmt.Print("a + b = %d", adder(a, b))
}


