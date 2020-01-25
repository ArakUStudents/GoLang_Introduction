package main

import "fmt"

func function() {
	println("'func function() {}' == 'void function() {}' in c family")
}

func argFunc(a int, b, c float32) {
	fmt.Printf("%d, %f, %f\n", a, b, c)
}

/*
func argFunc(string srt){
    go don't support overloading
}
*/

/*
func defaultArgValue(int x = 0){
	golang don't support default argument value
}
*/

func returnFunc() string {
	return ":/"
}

func int8MinMax() (int8, int8) {
	return -128, 127
}

func sum(numbers ...int) {
	fmt.Print(numbers, " ")
	total := 0
	for _, num := range numbers {
		total += num
	}
	fmt.Println(total)
}

func closureFunc() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func funcSample() {
	function()
	argFunc(1, 2, 3)
	println(returnFunc())
	min, max := int8MinMax()
	_, x := int8MinMax()
	y, _ := int8MinMax()
	min = y
	max = x
	println("int8 min:", min, "int8 max:", max)

	sum(1, 2)
	sum(123, 21321, 213213, 877, 65)
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum(numbers...)

	newInt := closureFunc()
	fmt.Println(newInt())
	fmt.Println(newInt())
	fmt.Println(newInt())
	fmt.Println(newInt())

	newInt = closureFunc()
	fmt.Println(newInt())
}
