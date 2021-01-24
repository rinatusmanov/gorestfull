package fiber_v2_driver

type IFiberCtx interface {
	OriginalURL() string
	Method(override ...string) string
	Body() (result []byte)
	Send(inData []byte) (err error)
}

type IOption interface{}