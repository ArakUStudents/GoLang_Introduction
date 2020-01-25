package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func goRoutineSample() {

	f("direct")

	go f("go routine1")
	go f("go routine2")
	go f("go routine3")

	time.Sleep(time.Second)
	fmt.Println("done")
}
