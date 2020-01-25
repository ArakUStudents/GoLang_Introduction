package main

import "fmt"

type (
	//Type is int
	Type = int
)

var inner bool = false

func pointerSample() {
	if inner {
		var pointerName *Type
		x := 123
		pointerName = &x
		simplePointer := &x
		fmt.Printf("%d == %d", *pointerName, *simplePointer)
	} else {
		inner = true
	}
	var functionPointer *func()
	(*functionPointer)()
}
