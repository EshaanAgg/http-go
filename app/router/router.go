package router

import (
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/parser"
)

func GetResponse(r *parser.Request) *parser.Response {
	// Route -> /
	if r.Target == "/" {
		return handleDefaultRoute(r)
	}

	// Route -> /echo/*
	if strings.HasPrefix(r.Target, "/echo/") {
		txt := r.Target[6:]
		return handleEchoRoute(r, txt)
	}

	// Route -> /user-agent
	if r.Target == "/user-agent" {
		return handleUserAgentRoute(r)
	}

	// Default route
	return handleNotFoundRoute(r)
}
