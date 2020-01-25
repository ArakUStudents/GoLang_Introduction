package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second * 5)
	fmt.Println("done")

	done <- true
}

func chanSyncSample() {

	done := make(chan bool, 1)
	go worker(done)
	<-done
}
