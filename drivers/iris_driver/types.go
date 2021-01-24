package iris_driver

import "net/http"

type irisCtx interface {
	Request() (request *http.Request)
	Write(body []byte) (n int, err error)
}

type IOption interface{}
