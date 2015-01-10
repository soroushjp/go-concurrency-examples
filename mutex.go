package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	state := make(map[int]int)
	mutex := &sync.Mutex{}
	var ops int64 = 0

	//10 concurrent readers to state
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)

				runtime.Gosched()
			}
		}()
	}

	//10 concurrent writers to state
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	//Let it run for 5 seconds
	time.Sleep(5 * time.Second)

	fmt.Println("Total operations: ", atomic.LoadInt64(&ops))

	mutex.Lock()
	fmt.Println("Final state: ", state)
	mutex.Unlock()
}
