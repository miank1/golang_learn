// what happens when a channel is closed and how the receiver detects that closure.

package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		defer close(ch)
		ch <- 10
		ch <- 20
		ch <- 30
	}()

	for x := range ch {
		fmt.Println(x)
	}

}
