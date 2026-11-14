package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex

	mu.Lock()

	fmt.Println("locked")

	mu.Lock()

	fmt.Println("done")
}
