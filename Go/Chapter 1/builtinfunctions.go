package main

import "fmt"


func main() {
  print_func()
  range_func()
}

func print_func(){
  fmt.Println("Hello")
  fmt.Print("Hello World")
}

func range_func(){
  for i :=1; i < 10; i++{
    fmt.Println(i)
  }
}
