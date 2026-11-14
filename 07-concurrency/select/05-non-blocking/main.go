package main

import "fmt"

func main() {
	ch := make(chan int)

	select {
	case value := <-ch:
		fmt.Println("received:", value)

	default:
		fmt.Println("channel not ready")
	}
}
