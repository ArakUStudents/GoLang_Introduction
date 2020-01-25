package main

import "fmt"

func forSample() {
	for {
		println("infinite loop")
		break
	}

	ch := "q"
	for {
		fmt.Scan(&ch)
		if ch != "q" {
			continue
		}
		break
	}

	println()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d, ", i+1)
	}
	println()

	arr := []string{"a", "b", "c", "d"}
	print("forr element:")
	for _, element := range arr {
		print(element, ", ")
	}
	println()

	print("forr index:")
	for index := range arr {
		print(index, ", ")
	}
	println()

	println()
	println("forr index & element:")
	for index, char := range arr {
		println("index:", index, "\tvalue:", char)
	}

	dictionary := map[string]string{"a": "apple", "b": "banana"}
	for key, value := range dictionary {
		fmt.Printf("%s -> %s\n", key, value)
	}

	println("forr string:")
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
