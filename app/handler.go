package main

import (
	"io"
	"log"
	"net"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/parser"
)

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Printf("[%s] Connection closed by the client\n", conn.RemoteAddr().String())
				return
			}

			log.Printf("[%s] Error reading message: %s\n", conn.LocalAddr(), err)
			return
		}

		r, err := parser.NewRequest(buf)
		if err != nil {
			log.Printf("Error in parsing the request buffer ('%s'): %s ", buf, err)
			return
		}

		log.Printf("[%s] %s request to %s", conn.RemoteAddr().String(), r.GetMethod(), r.Target)

		resp := s.getResponse(r)

		// If the request is a close request, add the Connection: close header to the same
		shouldClose := s.isCloseRequest(r)
		if shouldClose {
			resp.SetHeader("Connection", "close")
		}

		b := resp.GetBuffer()
		_, err = conn.Write(b.Bytes())
		if err != nil {
			log.Printf("[%s] Failed to send message (%s) : %v\n", conn.RemoteAddr().String(), b.Bytes(), err)
		}

		if shouldClose {
			log.Printf("[%s] Recieved a close request, terminating the connection", conn.RemoteAddr().String())
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
