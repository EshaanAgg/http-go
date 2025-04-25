package parser

type Response struct {
	statusCode int
	headers    map[string]string

	// Stores the bytes of the response that have been added by the appropiate writers
	buffer []byte
}

func NewResponse() *Response {
	return &Response{
		statusCode: 200,
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
