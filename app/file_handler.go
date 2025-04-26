package main

import "github.com/codecrafters-io/http-server-starter-go/app/parser"

func (s *Server) handleFile(req *parser.Request) *parser.Response {
	switch req.Method {
	case parser.GET:
		return s.handleGetFile(req)
	case parser.POST:
		return s.handlePostFile(req)
	default:
		return parser.NewResponse(405) // Method Not Allowed
	}
}

func (s *Server) handleGetFile(req *parser.Request) *parser.Response {
	fileName := req.Target[len("/files/"):]
	data, err := s.getFileContent(fileName)

	// Return 404 if the file is not found
	if err != nil {
		return parser.NewResponse(404)
	}

	// Return 200 with the file content
	return parser.NewOctetStreamResponse(200, data)
}

func (s *Server) handlePostFile(req *parser.Request) *parser.Response {
	return parser.NewResponse(405) // Method Not Allowed
}
