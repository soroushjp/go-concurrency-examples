package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	var ops uint64 = 0
	n := 50
	var t int64 = 1

	for i := 0; i < n; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Duration(t) * time.Second)

	opsFinal := atomic.LoadUint64(&ops) //Get ops safely while being used by goroutines

	fmt.Printf("ops after %d concurrent increments for %d seconds: %e\n", n, t, float64(opsFinal))
}
