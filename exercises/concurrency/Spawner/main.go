package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(10000)
	start := time.Now()
	for i := 0; i < 10000; i++ {
		go func() {
			defer wg.Done()
			//fmt.Println("Done")
		}()
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println(elapsed)

}
