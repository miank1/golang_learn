// Bidirectional Channels

package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		value := <-ch
		fmt.Println("Worker received:", value)
		value = value * 2
		ch <- value

	}()

	ch <- 5
	response := <-ch
	fmt.Println("Main received response:", response)

}
