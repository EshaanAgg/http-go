package main

import "github.com/codecrafters-io/http-server-starter-go/app/parser"

func (s *Server) handleFile(req *parser.Request) *parser.Response {
	fileName := req.Target[len("/files/"):]

	switch req.Method {
	case parser.GET:
		return s.handleGetFile(fileName)
	case parser.POST:
		return s.handlePostFile(req, fileName)
	default:
		return parser.NewResponse(405) // Method Not Allowed
	}
}

func (s *Server) handleGetFile(fileName string) *parser.Response {
	data, err := s.getFileContent(fileName)

	// Return 404 if the file is not found
	if err != nil {
		return parser.NewResponse(404)
	}

	// Return 200 with the file content
	return parser.NewOctetStreamResponse(200, data)
}

func (s *Server) handlePostFile(req *parser.Request, fileName string) *parser.Response {
	err := s.saveFileContent(fileName, req.GetBody())
	if err != nil {
		return parser.NewResponse(500)
	}

	return parser.NewResponse(201)
}
