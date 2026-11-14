// create an unbuffered integer channel

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("Sending Started ...")
		ch <- 200
		fmt.Println("Sending Completed ...")
	}()

	fmt.Println("Main Started")

	time.Sleep(2 * time.Second)
	value := <-ch

	fmt.Println("Sending Completed ...")
	fmt.Println("Value Received:", value)

}
