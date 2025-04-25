package router

import "github.com/codecrafters-io/http-server-starter-go/app/parser"

type defaultRoute struct{}

func (defaultRoute) Handle(req *parser.Request) *parser.Response {
	return parser.NewResponse(200)
}
