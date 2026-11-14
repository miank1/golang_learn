package main

import (
	"context"
	"fmt"
)

func main() {

	ctx := context.Background()
	fmt.Println(ctx)

	ctx1 := context.TODO()
	fmt.Println(ctx1)
}
