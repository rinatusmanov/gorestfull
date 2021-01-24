package gin_driver

import (
	"errors"
	"fmt"
	"github.com/rinatusmanov/gorestfull/drivers/http_driver"
	"github.com/rinatusmanov/gorestfull/types"
	"net/http"
	"reflect"
)

type driver struct {
	ginFuncAny           reflect.Value
	ginAnyTypeOfSliceOne reflect.Type
	ginFuncAnyResults    []reflect.Value
	makeFuncType         reflect.Type
	options              []IOption
}

func (d *driver) SetHandler(pattern string, fun types.THandler) (err error) {
	defer func() {
		if errRecover := recover(); errRecover != nil {
			err = errors.New(fmt.Sprint(errRecover))
		}
	}()

	fn := reflect.MakeFunc(d.makeFuncType, func(args []reflect.Value) (results []reflect.Value) {
		request := args[0].Elem().FieldByName("Request").Interface().(*http.Request)
		writer := args[0].Elem().FieldByName("Writer").Interface().(http.ResponseWriter)
		fun(http_driver.NewReadWriter(writer, request))
		return d.ginFuncAnyResults
	})

	slice := reflect.MakeSlice(d.ginAnyTypeOfSliceOne, 1, 1)
	slice.Index(0).Set(fn)

	d.ginFuncAny.CallSlice([]reflect.Value{reflect.ValueOf(pattern), slice})
	return
}

func NewDriver(app interface{}, options ...IOption) (result types.IHttpDriver) {
	val := reflect.ValueOf(app)
	fnType := val.MethodByName("Any")

	return &driver{
		ginFuncAny:           fnType,
		ginFuncAnyResults:    []reflect.Value{},
		options:              options,
		makeFuncType:         fnType.Type().In(1).Elem(),
		ginAnyTypeOfSliceOne: fnType.Type().In(1),
	}
}
