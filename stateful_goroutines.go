//Stateful go-routines
package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	var ops int64 = 0 //Operations counter

	//Set up read & write channels
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	//Set up a thread to handle access to state. Use of select allows state to be owned by one and only one thread at a time.
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp              //Use the response from our concurrent read response here
				atomic.AddInt64(&ops, 1) //Operations counter incremented
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp             //Use bool success of concurrent write here
				atomic.AddInt64(&ops, 1) //Operations counter incremented
			}
		}()
	}

	//fmt.Println("Press enter to quit ...")
	//var input string
	//fmt.Scanln(&input)
	time.Sleep(5 * time.Second)

	//Safely print total operations
	fmt.Println("Total operations completed:", atomic.LoadInt64(&ops))
	fmt.Println("Quitting ...")

}
