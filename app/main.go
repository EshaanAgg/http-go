package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/codecrafters-io/http-server-starter-go/app/parser"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		log.Fatalf("Error listening on port 4221: %v", err)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	r := parser.NewResponse()
	b := r.GetBuffer()
	_, err = conn.Write(b)
	if err != nil {
		fmt.Printf("Failed to send message to the connection: %s\n", b)
	}
}
