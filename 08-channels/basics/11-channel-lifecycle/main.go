package main

import "fmt"

func main() {
	ch := make(chan int)

	val := 10
	go func() {
		for i := 0; i < 3; i++ {
			ch <- val
			val += 10
		}
		close(ch)
	}()

	for x := range ch {
		fmt.Println(x)
	}

	fmt.Println("Channel Closed")

}
