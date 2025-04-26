package router

import "github.com/codecrafters-io/http-server-starter-go/app/parser"

func handleDefaultRoute(req *parser.Request) *parser.Response {
	return parser.NewResponse(200)
}

func handleNotFoundRoute(req *parser.Request) *parser.Response {
	return parser.NewResponse(404)
}

func handleEchoRoute(req *parser.Request, txt string) *parser.Response {
	return parser.NewPlainTextResponse(200, txt)
}

func handleUserAgentRoute(req *parser.Request) *parser.Response {
	user_agent, ok := req.Headers["User-Agent"]
	if !ok {
		user_agent = "Unknown"
	}
	return parser.NewPlainTextResponse(200, user_agent)
}
