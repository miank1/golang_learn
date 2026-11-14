package main

import (
	"context"
	"fmt"
)

func main() {

	parent := context.Background()

	ctx, cancel := context.WithCancel(parent)

	ctxA, cancelA := context.WithCancel(ctx)
	ctxB, cancelB := context.WithCancel(ctx)

	defer cancel()
	defer cancelA()
	defer cancelB()

	fmt.Println("Parent:", ctx.Err())
	fmt.Println("Child A:", ctxA.Err())
	fmt.Println("Child B:", ctxB.Err())

	cancelA()

	fmt.Println("Parent:", ctx.Err())
	fmt.Println("Child A:", ctxA.Err())
	fmt.Println("Child B:", ctxB.Err())

}
