package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(5)

	for i := 1; i <= 5; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Printf("Worker processed %d\n", i)
		}(i)
	}
	wg.Wait()
}
