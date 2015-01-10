package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(time.Second * 2)

	go func() {
		<-timer1.C
		fmt.Println("timer 1 expired")
	}()

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer 2 expired")
	}()

	if stop2 := timer2.Stop(); stop2 {
		fmt.Println("timer 2 stopped")
	}

	var input string
	fmt.Scanln(&input)

}
