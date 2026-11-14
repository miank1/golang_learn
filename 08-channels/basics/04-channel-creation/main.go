// Channel Creation and Nil vs Initialized Channels

package main

import "fmt"

func main() {
	var ch1 chan int
	ch2 := make(chan int)

	fmt.Println("Is ch1 nil?", ch1 == nil)
	fmt.Println("Is ch2 nil?", ch2 == nil)
}
