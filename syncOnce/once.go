package main

import (
	"fmt"
	"sync"
)

func main() {
	f := func() {
		fmt.Println("hello, world")
	}
	var once sync.Once

	for i := 0; i < 10; i++ {
		once.Do(f)
	}
}
