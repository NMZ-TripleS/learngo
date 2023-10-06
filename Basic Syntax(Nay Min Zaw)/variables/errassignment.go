package main

import (
	"fmt"
	"os"
)

func main() {
  filename := "myfile.txt"
  file, err := os.Open(filename)
  if err != nil {
    fmt.Println(err)
    return
  }
  
  fmt.Println(file)
}