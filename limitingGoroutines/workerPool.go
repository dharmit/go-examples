package main

import (
	"fmt"
	"sync"
	"time"
)

var concurrentMG = 4

func main() {
	var wg sync.WaitGroup
	podChan := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			// put the task to process on podChan
			podChan <- i
		}
		close(podChan)
	}()

	// spin up concurrentMG number of workers
	wg.Add(concurrentMG)
	for i := 0; i < concurrentMG; i++ {
		fmt.Printf("Spinning up a goroutine...\n")
		go func() {
			// read till podChan is empty
			for i := range podChan {
				fmt.Printf("%d\n", i)
				time.Sleep(1 * time.Second)
			}

			// no more pods to process
			wg.Done()
		}()
	}

	wg.Wait()
}
