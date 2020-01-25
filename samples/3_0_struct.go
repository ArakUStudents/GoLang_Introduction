package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func structSample() {

	fmt.Println(person{"Ali", 20})

	fmt.Println(person{name: "Farid", age: 30})

	fmt.Println(person{name: "Majid"})

	fmt.Println(&person{name: "Hasan", age: 40})

	fmt.Println(newPerson("Iraj"))

	s := person{name: "Saiid", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
