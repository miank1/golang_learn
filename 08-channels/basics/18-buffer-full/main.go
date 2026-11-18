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

	go func() {
		ch <- 30
	}()

	fmt.Println("Receive Value ", <-ch)

	go func() {
		wg.Wait()
		close(ch)
	}()

	fmt.Println("Length Again :", len(ch))

}
