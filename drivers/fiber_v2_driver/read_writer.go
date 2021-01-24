package fiber_v2_driver

import (
	"github.com/rinatusmanov/gorestfull/maker"
	"github.com/rinatusmanov/gorestfull/types"
	"net/url"
)

type readWriter struct {
	fiber    IFiberCtx
	query    types.IHttpDriverQuery
	response types.IResponse
	method   string
	body     []byte
}

func NewReadWriter(fiber IFiberCtx) (result types.IReadWriter) {
	url, _ := url.Parse(fiber.OriginalURL())
	rw := &readWriter{
		fiber:  fiber,
		query:  maker.Query(url),
		method: fiber.Method(),
		body:   fiber.Body(),
	}
	rw.response = maker.NewResponse(rw)
	result = rw
	return
}

func (r *readWriter) Method() (method string) {
	return r.method
}

func (r *readWriter) GetBody() (body []byte) {
	return r.body
}

func (r *readWriter) Write(body []byte) (count int, err error) {
	err = r.fiber.Send(body)
	if err == nil {
		count = len(body)
	}
	return
}

func (r *readWriter) Query() (query types.IHttpDriverQuery) {
	return r.query
}

func (r *readWriter) Response() types.IResponse {
	return r.response
}
