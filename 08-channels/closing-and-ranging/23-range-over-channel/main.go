package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		defer close(ch)
		ch <- 10
		ch <- 20
		ch <- 30
		ch <- 40
		ch <- 50
	}()

	for x := range ch {
		fmt.Println("Values :", x)
	}

	fmt.Println("All values received")

}
