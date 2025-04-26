package parser

import "fmt"

type Response struct {
	statusCode int
	headers    map[string]string
	body       string

	// Stores the bytes of the response that have been added by the appropiate writers
	buffer []byte
}

// Creates a new response with the provided status code.
func NewResponse(status int) *Response {
	return &Response{
		statusCode: status,
		headers:    make(map[string]string),
		buffer:     make([]byte, 0),
	}
}

// Creates a new response with the provided status code and body.
// The body is set to the provided string and the content type is set to text/plain.
func NewPlainTextResponse(statusCode int, body string) *Response {
	r := NewResponse(statusCode)

	r.SetHeader("Content-Type", "text/plain")
	r.SetHeader("Content-Length", fmt.Sprintf("%d", len(body)))
	r.SetBody(body)

	return r
}

func (r *Response) GetBuffer() []byte {
	r.writeStatusLine()
	r.writeHeaders()
	r.writeBody()

	return r.buffer
}

func (r *Response) SetHeader(header, value string) {
	r.headers[header] = value
}

func (r *Response) SetBody(body string) {
	r.body = body
}
