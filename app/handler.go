package main

import (
	"log"
	"net"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/parser"
)

func (s *Server) handleConnection(conn net.Conn) {
	for {
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

		log.Printf("[REQUEST] %s | %s | %s", conn.LocalAddr().String(), r.GetMethod(), r.Target)

		resp := s.getResponse(r)

		// If the request is a close request, add the Connection: close header to the same
		shouldClose := s.isCloseRequest(r)
		if shouldClose {
			r.Headers["Connection"] = "close"
		}

		b := resp.GetBuffer()
		_, err = conn.Write(b)
		if err != nil {
			log.Printf("Failed to send message (%s) to the connection: %s\n", b, err)
		}

		if shouldClose {
			log.Printf("Recieved a close request, terminating the connection %s", conn.LocalAddr().String())
			conn.Close()
			return
		}
	}
}

// Process the request and return the appropriate response
// The function checks for exact matches first, then checks for prefix matches
func (s *Server) getResponse(r *parser.Request) *parser.Response {
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

// Checks if the request is a "close" request.
// This is a special request with the header Connection, set to as "close"
func (s *Server) isCloseRequest(r *parser.Request) bool {
	val, ok := r.Headers["Connection"]
	return ok && val == "close"
}
