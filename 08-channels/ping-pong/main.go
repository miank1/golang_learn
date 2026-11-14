package main

import (
	"fmt"
	"time"
)

func ping(ch chan int) {
	for i := 0; i < 10; i++ {
		n := <-ch
		fmt.Println("Ping *: ", n)
		n++
		ch <- n
	}
}

func pong(ch chan int) {
	for i := 0; i < 10; i++ {
		n := <-ch
		fmt.Println("Pong +: ", n)
		n++
		ch <- n
	}
}

func main() {
	ch := make(chan int, 1)
	ch <- 1
	go ping(ch)
	go pong(ch)

	time.Sleep(3 * time.Second)

}
