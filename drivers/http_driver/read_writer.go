package http_driver

import (
	"github.com/gorestfull/maker"
	"github.com/rinatusmanov/gorestfull/types"
	"io/ioutil"
	"net/http"
)

type readWriter struct {
	writer   http.ResponseWriter
	request  *http.Request
	query    types.IHttpDriverQuery
	response types.IResponse
	method   string
	body     []byte
}

func NewReadWriter(writer http.ResponseWriter, request *http.Request) (result types.IReadWriter) {
	byteSlice, errReadFromRequest := ioutil.ReadAll(request.Body)
	if errReadFromRequest != nil {
		byteSlice = []byte{}
	}
	rw := &readWriter{
		writer:  writer,
		request: request,
		query:   maker.Query(request.URL),
		method:  request.Method,
		body:    byteSlice,
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
	return r.writer.Write(body)
}

func (r *readWriter) Query() (query types.IHttpDriverQuery) {
	return r.query
}

func (r *readWriter) Response() types.IResponse {
	return r.response
}
