// Receiving from a Closed Channel

package main

import "fmt"

func main() {

	ch := make(chan int, 2)
	ch <- 100
	ch <- 200

	close(ch)

	for i := 0; i < 3; i++ {
		value, ok := <-ch
		fmt.Printf("Value: %d, Open: %t\n", value, ok)
	}

}
