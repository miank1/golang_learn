// Understand how multiple goroutines communicate through a shared channel.

package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		ch <- 100
	}()

	go func() {
		defer wg.Done()
		ch <- 200
	}()

	go func() {
		defer wg.Done()
		ch <- 300
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	for x := range ch {
		fmt.Println(x)
	}

}
