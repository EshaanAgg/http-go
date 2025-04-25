package parser

type Response struct {
	StatusCode int

	buffer []byte
}

func NewResponse() *Response {
	return &Response{
		StatusCode: 200,
		buffer:     make([]byte, 0),
	}
}

func (r *Response) GetBuffer() []byte {
	return r.buffer
}
