package fiber_v2_driver

import (
	"errors"
	"fmt"
	"github.com/rinatusmanov/gorestfull/types"
	"reflect"
)

type driver struct {
	fiberFuncAll           reflect.Value
	fiberAllTypeOfSliceOne reflect.Type
	fiberFuncAllResults    []reflect.Value
	makeFuncType           reflect.Type
	options                []IOption
}

func (d *driver) SetHandler(pattern string, fun types.THandler) (err error) {
	defer func() {
		if errRecover := recover(); errRecover != nil {
			err = errors.New(fmt.Sprint(errRecover))
		}
	}()

	fn := reflect.MakeFunc(d.makeFuncType, func(args []reflect.Value) (results []reflect.Value) {
		fiber, ok := args[0].Interface().(IFiberCtx)
		if ok {
			fun(NewReadWriter(fiber))
		}
		return d.fiberFuncAllResults
	})

	slice := reflect.MakeSlice(d.fiberAllTypeOfSliceOne, 1, 1)
	slice.Index(0).Set(fn)

	d.fiberFuncAll.CallSlice([]reflect.Value{reflect.ValueOf(pattern), slice})
	return
}

func NewDriver(app interface{}, options ...IOption) (result types.IHttpDriver) {

	val := reflect.ValueOf(app)
	fnType := val.MethodByName("All")
	resultFunc := []reflect.Value{reflect.Zero(fnType.Type().In(1).Elem().Out(0))}

	return &driver{
		fiberFuncAll:           fnType,
		fiberFuncAllResults:    resultFunc,
		options:                options,
		makeFuncType:           fnType.Type().In(1).Elem(),
		fiberAllTypeOfSliceOne: fnType.Type().In(1),
	}
}
