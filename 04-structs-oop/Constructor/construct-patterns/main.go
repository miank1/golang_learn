package main

import (
	"fmt"
	"hello/04-structs-oop/Constructor/construct-patterns/server"
)

func main() {

	s := server.NewServer(8000)
	fmt.Println(s)
}
