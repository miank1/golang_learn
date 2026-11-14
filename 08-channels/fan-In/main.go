package main

import "fmt"

func mergeChannel(ch1, ch2, result chan int) <-chan int {

	for ch := range ch1 {
		result <- ch
	}

	for ch := range ch2 {
		result <- ch
	}

	close(result)
	return result
}

func main() {

	chan1 := make(chan int, 3)
	chan2 := make(chan int, 3)

	chan1 <- 1
	chan1 <- 2
	chan1 <- 3
	close(chan1)

	chan2 <- 4
	chan2 <- 5
	chan2 <- 6

	close(chan2)

	result := make(chan int, 6)

	mergeChannel(chan1, chan2, result)

	for x := range result {
		fmt.Println(x)
	}
}
