package main

import "fmt"

func arraySample() {
	var simpleArray [12]int
	simpleArray[0] = 1
	simpleArray[11] = 12
	fmt.Println(simpleArray)
	//  simpleArray[12] : out of range
	initArray := [3]float32{0.1, 0.2, 0.3}
	initArray[0]++
	fmt.Println(initArray)
	var twoDimensional [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDimensional[i][j] = i + j
		}
	}
	fmt.Println(twoDimensional)
}
