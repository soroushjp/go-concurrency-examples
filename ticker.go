package main

import (
	"fmt"
	"time"
)

func pinger() {
	fmt.Println("ping boomchakalaka")
}

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for range ticker.C {
			pinger()
		}
	}()

	go func() {
		time.Sleep(time.Second * 5)
		ticker.Stop()
		fmt.Println("Ticker stopped.")
	}()

	var input string
	fmt.Scanln(&input)
}
