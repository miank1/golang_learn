// how a buffered channel handles multiple values sent and received by goroutines.

package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 5)

	var wg sync.WaitGroup
	wg.Add(1)

	// sender
	go func() {
		defer wg.Done()
		ch <- 10
		ch <- 20
		ch <- 30
		ch <- 40
		ch <- 50
		close(ch)
	}()

	for x := range ch {
		fmt.Println("Value received ", x)
	}

	wg.Wait()
}
