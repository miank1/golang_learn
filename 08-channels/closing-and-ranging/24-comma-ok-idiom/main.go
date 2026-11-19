// Detect whether a channel is closed using the comma OK idiom.

package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	ch <- 10
	ch <- 20
	ch <- 30

	close(ch)

	for {
		value, ok := <-ch
		if !ok {
			fmt.Println("Channel is closed")
			break
		}

		fmt.Printf("Value: %d, Open: %t\n", value, ok)
	}
}
