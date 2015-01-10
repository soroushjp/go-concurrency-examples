package main

import (
	"fmt"
	"time"
)

func main() {

	// //--Basic rate limiting--

	// //Create requests channel
	// requests := make(chan int, 5)

	// //Simulate 5 requests arriving, close channel after
	// for i := 0; i < 5; i++ {
	// 	requests <- i
	// }
	// close(requests)

	// //Create a rate-limiting ticker than
	// limiter := time.Tick(time.Millisecond * 200)

	// for req := range requests {
	// 	<-limiter //Limited will only unblock every 200ms, rate-limiting the processing of requests
	// 	fmt.Println("Request:", req, "Time:", time.Now())
	// }

	//--Bursty limiter--

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	//Create burstyRequests channel and simulate 5 requests and then close
	burstyRequests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("Request:", req, "Time:", time.Now())
	}

}
