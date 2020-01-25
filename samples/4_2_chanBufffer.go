package main

import "fmt"

func chanBufferSample() {

	messages := make(chan string, 2)

	messages <- "msg1"
	messages <- "msg2"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
