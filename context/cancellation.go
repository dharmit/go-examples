package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// start a goroutine
	wg.Add(1)
	go routine(ctx)

	// sleep some time and trigger a cancel
	time.Sleep(5 * time.Second)
	cancel()
	// wait for go routines to stop
	wg.Wait()
}

func routine(ctx context.Context) {
	// a ticker to be used in for...select
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ctx.Done(): // if cancel function is called, stop running
			ticker.Stop()
			fmt.Println("received cancel")
			time.Sleep(2 * time.Second)
			wg.Done()
			return
		case <-ticker.C: // print a dummy message upon every tick
			fmt.Println("ticking...")
		}
	}
}
