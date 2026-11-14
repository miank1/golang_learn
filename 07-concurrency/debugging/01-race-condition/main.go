package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var mu sync.Mutex

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()

			counter++

		}()
	}

	wg.Wait()

	fmt.Println(counter)
}
