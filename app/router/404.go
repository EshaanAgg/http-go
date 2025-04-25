package router

import "github.com/codecrafters-io/http-server-starter-go/app/parser"

type notFoundRoute struct{}

func (notFoundRoute) Handle(req *parser.Request) *parser.Response {
	return parser.NewResponse(404)
}
