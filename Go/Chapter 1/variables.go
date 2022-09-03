package main

import (
	"fmt"
	"math"
)
func main() {
	variables1()
	variables2()
	math_func()
}

func variables1() {
	var dog_name string = "spot" //Variable declaration
	dog_name = "rex" //assigning variable to a new value
	dog_name = "t-" + dog_name
	fmt.Println(dog_name)
}

func variables2() {
	big := "large" //variable declaration
	big1 := 1000 * 1000
	fmt.Println(big, big1)
}

func math_func() {
	sum := 5+5
	sub := 5-5
	mul := 2*2
	div := 5/2
	intdiv := 5/2
  floordiv := math.Floor(5/2)
	exp := math.Pow(5,2)
	mod := 5%2 
	fmt.Println(sum, sub, mul, div, intdiv ,floordiv, exp, mod)
}

