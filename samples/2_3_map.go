package main

import "fmt"

func mapSample() {

	map1 := make(map[string]int)

	map1["key1"] = 7
	map1["key2"] = 13

	fmt.Println("map:", map1)

	value1 := map1["key1"]
	fmt.Println("value1: ", value1)

	fmt.Println("len:", len(map1))

	delete(map1, "key2")
	fmt.Println("map:", map1)

	_, prs := map1["key2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
