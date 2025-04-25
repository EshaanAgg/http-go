package main

import (
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		log.Fatalf("Error listening on port 4221: %v", err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v\n", err.Error())
		}
		go handleConnection(conn)
	}
}
