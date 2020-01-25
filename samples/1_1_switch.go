package main

import (
	"fmt"
	"time"
)

func getInt() int {
	var x int
	fmt.Scan(&x)
	return x
}

func switchSample() {
	switch time.Now().Weekday() {
	case time.Wednesday, time.Friday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	switch dayOfWeek := time.Now().Weekday(); dayOfWeek {
	case time.Wednesday, time.Friday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	//=========================================

	switch num := getInt(); {
	case num%2 == 0:
		println("even")
	case num%2 == 1:
		println("odd")
	}

	num := getInt()
	switch {
	case num%2 == 0:
		println("even")
	case num%2 == 1:
		println("odd")
	}

	//=========================================

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a Boolean")
		case int:
			fmt.Println("I'm an Integer")
		default:
			fmt.Printf("I Don't know my type, but compiler says I'm %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
