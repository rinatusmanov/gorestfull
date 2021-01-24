package maker

import (
	"encoding/json"
	"fmt"
	"github.com/rinatusmanov/gorestfull/types"
	"reflect"
)

type crud struct {
	db       types.IGormDB
	maxCount uint64
	driver   types.IHttpDriver
	err      error
}

func (c *crud) createMany(reflectType reflect.Type) reflect.Value {
	return reflect.
		New(reflect.SliceOf(reflectType))
}

func (c *crud) Crud(model interface{}, options ...types.Option) types.IMaker {
	val := reflect.ValueOf(model)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	c.logic(model, val.Type())
	return c
}

func (c *crud) logic(model interface{}, modelType reflect.Type, options ...types.Option) {
	hooks := NewHookStore(model)
	schema, errSchema := parse(c.db.GormDB(), model)
	if errSchema != nil {
		return
	}
	uri := "/crud/" + schema.Table()
	for _, option := range options {
		switch option.Type() {
		case uriSet:
			uri = fmt.Sprint(option.Value())
		}
	}
	c.err = c.driver.SetHandler(uri, func(rw types.IReadWriter) {
		var (
			result       = c.createMany(modelType)
			responseData = rw.Response().SetResult(result)
			tx           = c.db
			errHook      error
		)
		defer func() {
			if errHook == nil {
				_, _ = rw.Response().Write()
			}
		}()
		queryInterface := rw.Query()
		switch rw.Method() {
		case MethodGet:
			var (
				preloadSlice   []string
				preloadParam   = "preload"
				preloadDefault []string
				idSlice        []uint64
				idParam        = "id"
				idDefault      []uint64
				count          = 0
				countParam     = "Count"
				countDefault   = 100
				page           int
				pageParam      = "page"
				pageDefault    = 0
				unscoped       = false
				unscopedParam  = "unscoped"
			)
			queryInterface.
				Keys(preloadParam, idParam, countParam, pageParam, unscopedParam).
				Defaults(preloadDefault, idDefault, countDefault, pageDefault, false).
				Parse(&preloadSlice, &idSlice, &count, &page, &unscoped)
			tx = tx.
				Limit(count).
				Offset(page * count)
			if unscoped {
				tx = tx.Unscoped()
			}
			if len(idSlice) > 0 {
				tx = tx.Where("id in (?)", idSlice)
			}
			for _, preload := range preloadSlice {
				tx = tx.Preload(preload)
			}
			if tx, errHook = hooks.BeforeDBFind(responseData, tx, rw); errHook != nil {
				return
			}
			tx = tx.Find(result.Interface())
			responseData.Parse(result, tx)
			if tx, errHook = hooks.AfterDBFind(responseData, tx, rw); errHook != nil {
				return
			}
		case MethodDelete:
			var (
				idSlice       []uint64
				idParam       = "id"
				idDefault     []uint64
				unscoped      bool
				unscopedParam = "unscoped"
			)
			queryInterface.
				Keys(idParam, unscopedParam).
				Defaults(idDefault, false).
				Parse(&idSlice, &unscoped)
			if unscoped {
				tx = tx.Unscoped()
			}
			if tx, errHook = hooks.BeforeDBDelete(responseData, tx, rw); errHook != nil {
				return
			}
			tx = tx.Where("id in (?)", idSlice).Delete(model)
			responseData.Parse(reflect.ValueOf(idSlice), tx)
			if tx, errHook = hooks.AfterDBDelete(responseData, tx, rw); errHook != nil {
				return
			}
		case MethodPost:
			errUnmarshal := json.Unmarshal(rw.GetBody(), result.Interface())
			if errUnmarshal != nil {
				responseData.Error(errUnmarshal)
				return
			}
			if tx, errHook = hooks.BeforeDBCreate(responseData, tx, rw); errHook != nil {
				return
			}
			tx = tx.Create(result.Interface())
			responseData.Parse(result, tx)
			if tx, errHook = hooks.AfterDBCreate(responseData, tx, rw); errHook != nil {
				return
			}
		case MethodPut:
			errUnmarshal := json.Unmarshal(rw.GetBody(), result.Interface())
			if errUnmarshal != nil {
				responseData.Error(errUnmarshal)
				return
			}
			if tx, errHook = hooks.BeforeDBChange(responseData, tx, rw); errHook != nil {
				return
			}
			tx = tx.Save(result.Interface())
			responseData.Parse(result, tx)
			if tx, errHook = hooks.AfterDBChange(responseData, tx, rw); errHook != nil {
				return
			}
		}
	})
	fmt.Println(fmt.Sprintf("create crud methods on %s", uri))
	return
}

func (c *crud) Error() (err error) {
	return c.err
}

func Maker(driver types.IHttpDriver, dbInterface interface{}) (result types.IMaker, err error) {
	cr := &crud{
		driver: driver,
	}
	defer func() {
		result = cr
	}()
	gormDB, errDB := Db(dbInterface)
	if errDB != nil {
		cr.err = errDB
		return
	}
	cr.db = gormDB
	return
}
