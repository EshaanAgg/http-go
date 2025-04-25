package router

import "github.com/codecrafters-io/http-server-starter-go/app/parser"

type echoRoute struct {
	txt string
}

func newEchoRoute(txt string) *echoRoute {
	return &echoRoute{
		txt: txt,
	}
}

func (echoRoute) Handle(req *parser.Request) *parser.Response {
	// TODO: Implement the same correctly
	return parser.NewResponse(200)
}
