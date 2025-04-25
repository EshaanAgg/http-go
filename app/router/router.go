package router

import (
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/parser"
)

type Route interface {
	Handle(r *parser.Request) *parser.Response
}

func GetRoute(r *parser.Request) Route {
	// Route -> /
	if r.Target == "/" {
		return defaultRoute{}
	}

	// Route -> /echo/*
	if strings.HasPrefix(r.Target, "/echo/") {
		txt := r.Target[6:]
		return newEchoRoute(txt)
	}

	// Default route
	return notFoundRoute{}
}
