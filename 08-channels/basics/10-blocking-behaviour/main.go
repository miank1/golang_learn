// how an unbuffered channel blocks a sender until a receiver is ready.

package main

import "fmt"

func main() {

	ch := make(chan int)

	go func() {
		ch <- 100
		fmt.Println("Worker: send completed")
	}()

	fmt.Println("Main: waiting to receive")
	fmt.Println("Main: received ", <-ch)

}
