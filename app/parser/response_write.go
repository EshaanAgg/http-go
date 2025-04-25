package parser

import (
	"fmt"
	"log"
)

var statusMap = map[int]string{
	200: "OK",
	404: "Not Found",
}

func (r *Response) writeCRLF() {
	r.buffer = append(r.buffer, "\r\n"...)
}

func (r *Response) writeStatusLine() {
	status, ok := statusMap[r.statusCode]
	if !ok {
		log.Fatalf("Invalid status code (%d): No entry found for the same in statusMap", r.statusCode)
	}

	statusLine := fmt.Sprintf("HTTP/1.1 %d %s", r.statusCode, status)
	r.buffer = append(r.buffer, statusLine...)
	r.writeCRLF()
}

func (r *Response) writeHeaders() {
	for key, value := range r.headers {
		headerLine := fmt.Sprintf("%s: %s", key, value)
		r.buffer = append(r.buffer, headerLine...)
		r.writeCRLF() // Marks the end of the particular header
	}

	r.writeCRLF() // Marks the end of headers
}

func (r *Response) writeBody() {}
