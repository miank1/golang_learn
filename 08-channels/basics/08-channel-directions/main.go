// Implement send-only and receive-only channels in Go.

package main

import "fmt"

func main() {

	ch := make(chan int)

	go func() {
		sendData(ch)

	}()

	receiveData(ch)
}

func sendData(ch chan<- int) {
	ch <- 10
	ch <- 20
	ch <- 30

	close(ch)
}

func receiveData(ch <-chan int) {

	for x := range ch {
		fmt.Println(x)
	}

}
