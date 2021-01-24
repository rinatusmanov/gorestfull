package maker

import "github.com/rinatusmanov/gorestfull/types"

const (
	// строковое значение метода http.Get
	MethodGet = "GET"
	// строковое значение метода http.Head
	MethodHead = "HEAD"
	// строковое значение метода http.Post
	MethodPost = "POST"
	// строковое значение метода http.Put
	MethodPut = "PUT"
	// строковое значение метода http.Patch
	MethodPatch = "PATCH" // RFC 5789
	// строковое значение метода http.Delete
	MethodDelete = "DELETE"
	// строковое значение метода http.Connect
	MethodConnect = "CONNECT"
	// строковое значение метода http.Options
	MethodOptions = "OPTIONS"
	// строковое значение метода http.Trace
	MethodTrace = "TRACE"
)

// опции
const (
	uriSet types.TInitOptionType = "URI_SET" // смена ендпоинта
)