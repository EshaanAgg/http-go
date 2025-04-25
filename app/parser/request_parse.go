package parser

import "fmt"

// Parses the request buffer to extract the next word.
// The parsing is stopped when either a space of a CRLF is encountered.
func (r *Request) parseNextWord() string {
	w := make([]byte, 0)
	for r.buf[r.idx] != ' ' && r.buf[r.idx] != '\r' {
		w = append(w, r.buf[r.idx])
		r.idx++
	}

	if r.buf[r.idx] == ' ' {
		r.idx++
	} else {
		// Ending with CRLF
		r.idx += 2
	}

	return string(w)
}

func (r *Request) consumeCLRF() {
	r.idx += 2
}

// Parses a sample request line for the provided request.
// Eg. GET /index.html HTTP/1.1\r\n
func (r *Request) parseRequestLine() error {
	err := r.parseHTTPMethod()
	if err != nil {
		return err
	}

	r.Target = r.parseNextWord()
	r.parseNextWord() // TODO: Handle parsing of HTTP/1.1 and reject other requests
	r.consumeCLRF()

	return nil
}

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
