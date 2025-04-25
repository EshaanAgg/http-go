package parser

type Response struct {
	statusCode int
	headers    map[string]string
	body       string

	// Stores the bytes of the response that have been added by the appropiate writers
	buffer []byte
}

func NewResponse(status int) *Response {
	return &Response{
		statusCode: status,
		headers:    make(map[string]string),
		buffer:     make([]byte, 0),
	}
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
