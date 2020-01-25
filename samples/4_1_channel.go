package main

import "fmt"

func channelSample() {

	messages := make(chan string)

	go func() { messages <- "message" }()

	msg := <-messages
	fmt.Println(msg)
	close(messages)
}
