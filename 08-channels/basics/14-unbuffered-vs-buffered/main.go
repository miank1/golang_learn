// Write a Go program to demonstrate the difference between buffered and unbuffered channels.

package main

import "fmt"

func main() {

	// buffered channel
	ch := make(chan int, 2)
	ch <- 10
	ch <- 20

	fmt.Println("length & Capacity : ", len(ch), cap(ch))

	close(ch)

	// unbuffered channel
	ch1 := make(chan int)
	go func() {
		ch1 <- 100
	}()

	fmt.Println(<-ch1)

}
