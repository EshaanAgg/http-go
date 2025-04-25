package parser

type HTTPMethod int

const (
	GET HTTPMethod = iota
	POST
)

type Request struct {
	Method  HTTPMethod
	Target  string
	Version float64
	Headers map[string]string

	buf []byte
	idx int
}

func NewRequest(buf []byte) (*Request, error) {
	r := Request{
		buf: buf,
		idx: 0,
	}

	err := r.parseRequestLine()
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
