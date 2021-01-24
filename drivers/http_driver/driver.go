package http_driver

import (
	"errors"
	"fmt"
	"github.com/rinatusmanov/gorestfull/types"
	"net/http"
)

type driver struct {
	mux     *http.ServeMux
	options []IOption
}

func (d *driver) SetHandler(pattern string, fun types.THandler) (err error) {
	defer func() {
		if errRecover := recover(); errRecover != nil {
			err = errors.New(fmt.Sprint(errRecover))
		}
	}()

	d.mux.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
		fun(NewReadWriter(writer, request))
	})
	return
}

func NewDriver(mux *http.ServeMux, options ...IOption) types.IHttpDriver {
	return &driver{
		mux:     mux,
		options: options,
	}
}
