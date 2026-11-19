// what happens when a channel is closed and how the receiver detects that closure.

package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		defer close(ch)
		ch <- 10
		ch <- 20
		ch <- 30
	}()

	value1 := <-ch
	fmt.Println("Received:", value1)

	value2 := <-ch
	fmt.Println("Received:", value2)

	value3 := <-ch
	fmt.Println("Received:", value3)

	_, ok := <-ch
	if !ok {
		fmt.Println("Channel closed")
	}

}
