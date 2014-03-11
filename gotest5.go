package main

import "fmt"

const(
  MAX int = 10
)

func main() {
  var sum int;

  for i := 0; i < MAX; i ++ {
    sum += i 
  }

  fmt.Print(sum)
  
}

