package main

import "fmt"

func main() {

	ch := make(chan int, 5)
	ch <- 10
	ch <- 20
	ch <- 30

	fmt.Println("Length & Capacity", len(ch), cap(ch))

	ch <- 40
	ch <- 50

	fmt.Println("Length & Capacity", len(ch), cap(ch))

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("Length & Capacity after receiving ", len(ch), cap(ch))

}
