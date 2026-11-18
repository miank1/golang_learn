// Understand what happens when a receiver tries to read from an empty buffered channel.

package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan int, 2)
	var wg sync.WaitGroup
	wg.Add(1)

	// receiver
	go func() {
		defer wg.Done()
		fmt.Println("Receiver: waiting for value ")
		value := <-ch
		fmt.Println("Receiver : received ", value)
	}()

	ch <- 100

	fmt.Println("Main: sent ", <-ch)
	wg.Wait()

}
