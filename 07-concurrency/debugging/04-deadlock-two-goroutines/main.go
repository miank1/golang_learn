package main

import (
	"fmt"
	"sync"
)

func main() {

	var (
		mu1 sync.Mutex
		mu2 sync.Mutex
	)

	go func() {
		mu1.Lock()
		defer mu1.Unlock()

		mu2.Lock()
		defer mu2.Unlock()

		fmt.Println("worker 1")
	}()

	go func() {
		mu2.Lock()
		defer mu2.Unlock()

		mu1.Lock()
		defer mu1.Unlock()

		fmt.Println("worker 2")
	}()

}
