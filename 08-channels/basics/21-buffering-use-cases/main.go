// Simulate a small job queue using a buffered channel.

package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan string, 3)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		ch <- "Send Email"
		ch <- "Generate Report"
		ch <- "Process Payment"
		close(ch)
		fmt.Println("Channel Length & Capacity ", len(ch), cap(ch))

	}()

	go func() {
		defer wg.Done()
		for x := range ch {
			fmt.Println("Processing ... '", x)
		}
	}()

	wg.Wait()

}
