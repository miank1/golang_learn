// Channel Direction

package main

import "fmt"

func main() {
	var ch1 chan int
	var ch2 chan string
	var ch3 chan bool

	ch1 = make(chan int)
	ch2 = make(chan string)
	ch3 = make(chan bool)

	go func() {
		ch1 <- 20
		ch2 <- "Hello World"
		ch3 <- true
	}()

	fmt.Println("Recived Values ", <-ch1, <-ch2, <-ch3)
}
