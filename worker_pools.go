package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {

	//Create jobs channel to send jobs to workers
	jobs := make(chan int, 100)
	//Create results channel to receive results from workers
	results := make(chan int, 100)

	//Spawn 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	//Send over 9 jobs
	for j := 1; j <= 9; j++ {
		jobs <- j

	}
	//Close jobs channel when we have sent all over
	close(jobs)

	//Print results from worker
	for a := 1; a <= 9; a++ {
		fmt.Println("Result: ", <-results)
	}

}
