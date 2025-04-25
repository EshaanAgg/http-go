package parser

import "fmt"

func (r *Response) WriteOk() {
	r.buffer = append(r.buffer, "OK"...)
}

func (r *Response) WriteCRLF() {
	r.buffer = append(r.buffer, "\r\n"...)
}

func (r *Response) WriteHeader() {
	header := fmt.Sprintf("HTTP/1.1 %d ", r.StatusCode)
	r.buffer = append(r.buffer, []byte(header)...)
}
