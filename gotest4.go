package main

import "fmt"

type CustomPackge struct{
  age int
  name string
}

func main() {
  var fuck CustomPackge

  fuck.age = 10
  fuck.name = "fuck you"
  fmt.Print(fuck.age, fuck.name)
}

