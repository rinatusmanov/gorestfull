//

package maker

import (
	"fmt"
	"github.com/rinatusmanov/gorestfull/types"
	"reflect"
)

func parse(db interface{}, model interface{}) (schema types.IGormSchema, err error) {
	//defer func() {
	//	if panicInterface := recover(); panicInterface != nil {
	//		err = errors.New(fmt.Sprint(panicInterface))
	//	}
	//}()
	val := reflect.ValueOf(db)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	Statement := val.FieldByName("Statement")
	Parse := Statement.MethodByName("Parse")
	Parse.Call([]reflect.Value{reflect.ValueOf(model)})
	if Statement.Kind() == reflect.Ptr {
		Statement = Statement.Elem()
	}
	Schema := Statement.FieldByName("Schema")
	if Schema.Interface() != nil && Schema.Type().String() == "*schema.Schema" {
		schema = &gormSchema{
			value: Schema.Elem(),
			table: fmt.Sprint(Schema.Elem().FieldByName("Table").Interface()),
		}
	}
	return
}

type gormSchema struct {
	value reflect.Value
	table string
}

func (g *gormSchema) Table() string {
	return g.table
}
