package main

import (
	"log"
	"net"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/parser"
)

func (s *Server) handleConnection(conn net.Conn) {
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

	resp := s.GetResponse(r)
	b := resp.GetBuffer()
	_, err = conn.Write(b)
	if err != nil {
		log.Printf("Failed to send message (%s) to the connection: %s\n", b, err)
	}
}

func (s *Server) GetResponse(r *parser.Request) *parser.Response {
	var exactMatchRoutes = map[string]func(*parser.Request) *parser.Response{
		"/":           s.handleDefault,
		"/user-agent": s.handleUserAgent,
	}
	var prefixMatchRoutes = map[string]func(*parser.Request) *parser.Response{
		"/echo/":  s.handleEcho,
		"/files/": s.handleFile,
	}

	// Check for exact matches first
	if handler, ok := exactMatchRoutes[r.Target]; ok {
		return handler(r)
	}

	// Check for prefix matches
	for prefix, handler := range prefixMatchRoutes {
		if strings.HasPrefix(r.Target, prefix) {
			return handler(r)
		}
	}

	return s.handleNotFound(r)
}
