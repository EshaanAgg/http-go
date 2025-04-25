package main

import (
	"log"
	"net"

	"github.com/codecrafters-io/http-server-starter-go/app/parser"
	"github.com/codecrafters-io/http-server-starter-go/app/router"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Printf("error reading message from the connection %s: %s\n", conn.LocalAddr(), err)
		return
	}

	r, err := parser.NewRequest(buf)
	if err != nil {
		log.Printf("error in parsing the request buffer ('%s'): %s ", buf, err)
		return
	}

	log.Printf("[INFO] %s request to '%s'", r.GetMethod(), r.Target)
	hdlr := router.GetRoute(r)

	resp := hdlr.Handle(r)
	b := resp.GetBuffer()
	_, err = conn.Write(b)
	if err != nil {
		log.Printf("Failed to send message (%s) to the connection: %s\n", b, err)
	}
}
