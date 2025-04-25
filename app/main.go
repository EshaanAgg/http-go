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

	buf := make([]byte, 1024)
	conn.Read(buf)
	r, err := parser.NewRequest(buf)
	if err != nil {
		fmt.Printf("There was an error in parsing the request buffer ('%s'): %s ", buf, err)
		return
	}

	if r.Target == "/" {
		resp := parser.NewResponse(200)
		b := resp.GetBuffer()
		_, err = conn.Write(b)
		if err != nil {
			fmt.Printf("Failed to send message to the connection: %s\n", b)
		}
	} else {
		resp := parser.NewResponse(404)
		b := resp.GetBuffer()
		_, err = conn.Write(b)
		if err != nil {
			fmt.Printf("Failed to send message to the connection: %s\n", b)
		}
	}
}
