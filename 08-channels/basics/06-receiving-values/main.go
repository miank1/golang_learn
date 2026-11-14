// Receive and Process Values

package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 10
		ch <- 20
		ch <- 30
	}()

	first := <-ch
	second := <-ch
	third := <-ch
	fmt.Println("result ", first+second+third)

}
