package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{}

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			numbers = append(numbers, i)
		}(i)
	}

	wg.Wait()

	fmt.Println(numbers)
}
