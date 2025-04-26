package parser

import (
	"fmt"
	"strconv"
	"strings"
)

// Parses the request buffer until ine of the specified delimiters is found.
func (r *Request) parseUntil(cutset string) string {
	w := make([]byte, 0)

	for !strings.Contains(cutset, string(r.buf[r.idx])) {
		w = append(w, r.buf[r.idx])
		r.idx++
		if r.idx >= len(r.buf) {
			return string(w)
		}
	}

	r.idx++ // Skip the delimiter
	return string(w)
}

// Parses the request buffer to extract the next word.
// The parsing is stopped when either a space of a CRLF is encountered.
func (r *Request) parseNextWord() string {
	s := r.parseUntil(" \r")

	// Skip the end of CRLF if present
	if r.idx < len(r.buf) && r.buf[r.idx] == '\n' {
		r.idx++
	}

	return s
}

func (r *Request) consumeCLRF() {
	r.idx += 2
}

func (r *Request) parse() error {
	err := r.parseRequestLine()
	if err != nil {
		return err
	}

	err = r.parseHeaders()
	if err != nil {
		return err
	}

	if contentLength, ok := r.Headers["Content-Length"]; ok {
		r.parseBody(contentLength)
	}

	return nil
}

// Parses a sample request line for the provided request.
// Eg. GET /index.html HTTP/1.1\r\n
func (r *Request) parseRequestLine() error {
	err := r.parseHTTPMethod()
	if err != nil {
		return err
	}

	r.Target = r.parseNextWord()
	r.parseHTTPVersion()
	r.consumeCLRF()

	return nil
}

// Parses the HTTP version of the request.
// For eg. HTTP/1.1
// Also asserts the version to be 1.1 as the server only supports HTTP/1.1 currently.
func (r *Request) parseHTTPVersion() error {
	version := r.parseNextWord()
	if version != "HTTP/1.1" {
		return fmt.Errorf("invalid HTTP version found for the request: %s", version)
	}

	return nil
}

// Parses the headers of the request.
// The parsing is stopped when a CRLF is encountered.
func (r *Request) parseHeaders() error {
	for r.buf[r.idx] != '\r' {
		key := r.parseUntil(":")
		if key == "" {
			return fmt.Errorf("invalid header key found for the request: %s", key)
		}
		value := r.parseUntil("\r")
		if value == "" {
			return fmt.Errorf("invalid header value found for the request: %s", value)
		}
		r.Headers[key] = strings.TrimLeft(value, " ")
		r.idx++ // Skip the end of CRLF
	}

	r.consumeCLRF()
	return nil
}

// Parses the HTTP method of the request, and sets the Method field of the request.
func (r *Request) parseHTTPMethod() error {
	m := r.parseNextWord()
	if m == "GET" {
		r.Method = GET
	} else if m == "POST" {
		r.Method = POST
	} else {
		return fmt.Errorf("invalid HTTP method found for the request: %s", m)
	}

	return nil
}

// Parses the body of the request.
func (r *Request) parseBody(contentLength string) {
	cl, err := strconv.Atoi(contentLength)
	if err != nil {
		fmt.Printf("error in parsing the content length: %s", err)
		return
	}

	r.body = r.buf[r.idx : r.idx+cl]
	r.idx += cl
}
