// Understand how a buffered channel stores values without requiring an immediate receiver.

package main

import "fmt"

func main() {

	ch := make(chan int, 3)

	fmt.Println("Length of channel", len(ch))
	fmt.Println("Channel capacity", cap(ch))
	ch <- 10
	ch <- 20
	ch <- 30
	close(ch)
	fmt.Println("Length of channel", len(ch))
	fmt.Println("Channel capacity", cap(ch))

	for c := range ch {
		fmt.Println(c)
	}

}
