// when you attempt to send more values than a buffered channel can hold.

package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 2)
	var wg sync.WaitGroup

	ch <- 10
	ch <- 20

	fmt.Println("Initial Length :", len(ch))
	wg.Add(1)

	// sender
	go func() {
		defer wg.Done()
		ch <- 30
	}()

	// receiver
	fmt.Println("Receive Value ", <-ch)
	wg.Wait()
	fmt.Println("Length Again :", len(ch))

}
