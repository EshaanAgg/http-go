package router

import (
	"fmt"

	"github.com/codecrafters-io/http-server-starter-go/app/parser"
)

type echoRoute struct {
	txt string
}

func newEchoRoute(txt string) *echoRoute {
	return &echoRoute{
		txt: txt,
	}
}

func (e *echoRoute) Handle(req *parser.Request) *parser.Response {
	r := parser.NewResponse(200)

	r.SetHeader("Content-Type", "text/plain")
	r.SetHeader("Content-Length", fmt.Sprintf("%d", len(e.txt)))
	r.SetBody(e.txt)

	return r
}
