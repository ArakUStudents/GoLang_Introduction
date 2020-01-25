package main

import (
	"fmt"
)

type point struct {
	x, y int
}

func printSample() {
	print("hi")
	println()
	println("this is a new line")
	fmt.Print("this is another print func")
	fmt.Println()
	fmt.Printf("this is a printf of number %d", 1234)
}
