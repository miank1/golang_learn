package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("Working")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	cancel()

	fmt.Println("Stopped")
	<-done
}
