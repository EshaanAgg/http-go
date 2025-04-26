package main

import (
	"log"
	"net"
)

type Server struct {
	Port         int
	FilesBaseDir string
}

func NewServer(port int, filesBaseDir string) *Server {
	return &Server{
		Port:         port,
		FilesBaseDir: filesBaseDir,
	}
}

func (s *Server) Start() {
	log.Printf("Starting server on port %d\n", s.Port)
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		log.Fatalf("Error listening on port 4221: %v", err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v\n", err.Error())
		}
		go s.handleConnection(conn)
	}
}
