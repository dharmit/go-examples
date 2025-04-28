package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// number of concurrent goroutines we want to run
const workers = 4

func main() {
	var wg sync.WaitGroup
	// channel to be populated and depopulated at the start and end of each goroutine
	w := make(chan struct{}, workers)

	for i := 0; i < 10; i++ {
		w <- struct{}{}
		wg.Add(1)

		go func(x int) {
			defer func() {
				<-w
				wg.Done()
			}()

			fmt.Println("# of goroutines:", runtime.NumGoroutine(), "; goroutine #", x)
			time.Sleep(5 * time.Second)
		}(i)
	}
	wg.Wait()
}
