// Understand the behaviour of nil channels.

package main

import "fmt"

func main() {
	var ch1 chan int
	fmt.Println("ch2 is nil", ch1 == nil)

	ch2 := make(chan int)
	fmt.Println("ch2 is nil", ch2 == nil)

}
