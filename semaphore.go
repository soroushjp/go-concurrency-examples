// Semaphores

package main

import (
	"fmt"
	"time"
)

func Worker(j int, sem chan bool, results chan int) {
	sem <- true

	fmt.Println("Semaphore allowed access, beginning work ...")
	time.Sleep(1 * time.Second)
	results <- j * 2

	<-sem
}

func main() {
	sem := make(chan bool, 5) // Create buffered channel that acts as semaphore allowing 5 concurrent accesses
	results := make(chan int) // Setup results channel to get back results from workers

	// Create 10 workers
	for i := 0; i < 10; i++ {
		go Worker(i, sem, results)
	}

	// Wait for 10 results
	for i := 0; i < 10; i++ {
		fmt.Println(<-results)
	}

	// Keep main thread alive so goroutines can finish until user presses Enter key
	fmt.Println("Press enter key to quit ...")
	var dummyInput string
	fmt.Scanln(&dummyInput)
}
