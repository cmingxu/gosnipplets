package main

import "fmt"
import "os"

func main() {
  file, err := os.Open("gotest6.go")
  if err != nil {
    fmt.Print(err)
  }
  defer file.Close()

  data := make([]byte, 200)
  count, err = file.Read(data)

  if err != nil {
    fmt.Print(err)
  }

  fmt.Printf("read %d bytes: %q\n", count, data[:count])

}

