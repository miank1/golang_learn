package main

import "fmt"

var counter = 1

func main() {

	for i := 1; i <= 10000; i++ {
		go func() {
			counter++
		}()
	}

	fmt.Println("Counter Value ", counter)

}
