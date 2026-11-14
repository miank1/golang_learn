package main

import "fmt"

func main() {

	ch1 := make(chan string)
	ch1 <- "user service"

	ch2 := make(chan string)
	ch2 <- "payment service"

	go func() {
		ch1 <- "user service"
	}()

	go func() {
		ch2 <- "payment service"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println(msg)

		case msg := <-ch2:
			fmt.Println(msg)
		}
	}

}
