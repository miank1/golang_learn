package server

import "time"

type server struct { // private struct
	port    int
	host    string
	timeout time.Duration
}

// constructor pattern

func NewServer(port int) *server { // exported

	return &server{
		port:    port,
		host:    "127.80.102.1",
		timeout: 30 * time.Second,
	}
}
