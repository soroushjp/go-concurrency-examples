package main

import (
	"fmt"
	"time"
)

func pinger(c chan<- string) {
	for {
		c <- "ping"
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan<- string) {
	for {
		c <- "pong"
		time.Sleep(time.Second * 2)
	}
}

func printer(c1, c2 <-chan string) {
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		case <-time.After(time.Second * 10):
			fmt.Println("Timeout, printer exited.")
			return
		default:
			fmt.Println("do something else here while we wait. Non-blocking selects!")
			time.Sleep(time.Second * 1)
		}
	}
}

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go pinger(c1)
	go ponger(c2)
	go printer(c1, c2)

	var inputStr string
	fmt.Scanln(&inputStr)

}
