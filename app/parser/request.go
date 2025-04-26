package parser

import "strings"

type HTTPMethod int

const (
	GET HTTPMethod = iota
	POST
)

type Request struct {
	Method  HTTPMethod
	Target  string
	Headers map[string]string

	// Bytes in the body of the request. It is set
	// after the parsing of the request is done.
	body []byte

	buf []byte // Bytes of the request
	idx int
}

func NewRequest(buf []byte) (*Request, error) {
	r := Request{
		buf:     buf,
		idx:     0,
		Headers: make(map[string]string),
	}

	err := r.parse()
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (r *Request) GetMethod() string {
	switch r.Method {
	case GET:
		return "GET"
	case POST:
		return "POST"
	default:
		return "UNRECOGNIZED_HTTP_VERB"
	}
}

func (r *Request) GetBody() []byte {
	return r.body
}

func (r *Request) GetEncoding() SupportedEncodingType {
	encodings, ok := r.Headers["Accept-Encoding"]
	if !ok {
		return NoEncoding
	}

	supportedEncodings := strings.SplitSeq(encodings, ",")
	for encoding := range supportedEncodings {
		switch strings.TrimSpace(encoding) {
		// Arrange the supported encodings in the order of preference
		case "gzip":
			return Gzip
		case "identity":
			return NoEncoding
		}
	}

	// Default to no encoding if none of the supported encodings are found
	return NoEncoding
}
