package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	stop := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Pulse")
			case <-stop:
				fmt.Println("Stop")
				return
			}
		}
	}()

	time.Sleep(3 * time.Second)

	close(stop)

	time.Sleep(100 * time.Millisecond)

}
