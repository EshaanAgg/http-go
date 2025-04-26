package main

import "github.com/codecrafters-io/http-server-starter-go/app/parser"

func (s *Server) handleDefault(req *parser.Request) *parser.Response {
	return parser.NewResponse(200)
}

func (s *Server) handleNotFound(*parser.Request) *parser.Response {
	return parser.NewResponse(404)
}

func (s *Server) handleEcho(req *parser.Request) *parser.Response {
	txt := req.Target[len("/echo/"):]
	return parser.NewPlainTextResponse(200, txt)
}

func (s *Server) handleUserAgent(req *parser.Request) *parser.Response {
	user_agent, ok := req.Headers["User-Agent"]
	if !ok {
		user_agent = "Unknown"
	}
	return parser.NewPlainTextResponse(200, user_agent)
}
