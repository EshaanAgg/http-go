package parser

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"log"
)

type SupportedEncodingType int

const (
	NoEncoding SupportedEncodingType = iota
	Gzip
)

type Response struct {
	statusCode int
	headers    map[string]string
	// The content in the body of the response.
	body []byte

	// Stores the bytes of the response that have been added by the appropiate writers
	buffer bytes.Buffer
}

// Creates a new response with the provided status code.
func NewResponse(status int) *Response {
	return &Response{
		statusCode: status,
		headers:    make(map[string]string),
	}
}

// Encodes the response body based on the provided encoding type.
// If a supported encoding type is provided, the appropriate encoding is applied to the body, and
// the "Content-Encoding" header is set accordingly.
func getEncodedBody(r *Response, encoding SupportedEncodingType, body []byte) []byte {
	switch encoding {
	case NoEncoding:
		return body

	case Gzip:
		r.SetHeader("Content-Encoding", "gzip")

		var compressedBody bytes.Buffer
		w := gzip.NewWriter(&compressedBody)
		if _, err := w.Write([]byte(body)); err != nil {
			log.Fatalf("failed to write gzip: %v", err)
		}
		if err := w.Close(); err != nil {
			log.Fatalf("failed to close gzip writer: %v", err)
		}
		return compressedBody.Bytes()
	}

	panic(fmt.Sprintf("Unsupported encoding type: %d", encoding))
}

// Creates a new response with the provided status code and body.
// The body is set to the provided string and the content type is set to text/plain.
func NewPlainTextResponse(statusCode int, body []byte, encoding SupportedEncodingType) *Response {
	r := NewResponse(statusCode)

	body = getEncodedBody(r, encoding, body)
	r.SetHeader("Content-Type", "text/plain")
	r.SetHeader("Content-Length", fmt.Sprintf("%d", len(body)))
	r.SetBody(body)

	return r
}

func NewOctetStreamResponse(statusCode int, body []byte) *Response {
	r := NewResponse(statusCode)

	r.SetHeader("Content-Type", "application/octet-stream")
	r.SetHeader("Content-Length", fmt.Sprintf("%d", len(body)))
	r.SetBody(body)

	return r
}

func (r *Response) GetBuffer() bytes.Buffer {
	r.writeStatusLine()
	r.writeHeaders()
	r.writeBody()

	return r.buffer
}

func (r *Response) SetHeader(header, value string) {
	r.headers[header] = value
}

func (r *Response) SetBody(body []byte) {
	r.body = body
}
