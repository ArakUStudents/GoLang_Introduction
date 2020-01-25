package main

import "fmt"

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func methodSample() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perimeter:", r.perimeter())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perimeter:", rp.perimeter())
}
