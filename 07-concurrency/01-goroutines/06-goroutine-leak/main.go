package main

import (
	"fmt"
	"sync"
)

func worker(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range ch {
		fmt.Println("processing", job)
	}
}

func main() {
	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)

	go worker(ch, &wg)

	ch <- 10

	fmt.Println("main waiting")

	close(ch)
	wg.Wait()

	fmt.Println("main done")
}
