package main

import (
	"fmt"
	"math/rand"
)


func main() {
  if_else()
  for_loop()
  while_loop()
  switch_func()
}

func if_else(){
  i := rand.Intn(50)
  if i == 45 {
    fmt.Println("i is 45")
  } else if i == 40{
    fmt.Println("i is 40")
  } else {
    fmt.Println("i is", i)
  }
}

//Golang's for loop is similar to C
func for_loop(){
  for i := 1; i<20; i++{
    fmt.Println(i)
  }
}

//Golang doesn't have a while loop. Use for instead
func while_loop(){
  count := 0
  for count < 3{
    fmt.Println("The count is", count)
    count++
  }
}

//Golang also has Switch like in C
func switch_func(){
  switch i := rand.Intn(50); i{
    case 45:
      fmt.Println("The number is 45")
    case 40:
      fmt.Println("The number is 40")
    default:
      fmt.Println("The number is", i)
  }
  
}