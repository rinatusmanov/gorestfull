package maker

import (
	"errors"
	"github.com/rinatusmanov/gorestfull/types"
	"reflect"
)

type db struct {
	gormDB         reflect.Value // pointer to *gorm.DB
	methodLimit    reflect.Value // reflect.Value of method Limit of *grom.DB getted from *gorm.DB type
	methodOffset   reflect.Value // reflect.Value of method Offset of *grom.DB getted from *gorm.DB type
	methodCount    reflect.Value // reflect.Value of method Count of *grom.DB getted from *gorm.DB type
	methodUnscoped reflect.Value // reflect.Value of method Unscoped of *grom.DB getted from *gorm.DB type
	methodCreate   reflect.Value // reflect.Value of method Create of *grom.DB getted from *gorm.DB type
	methodSave     reflect.Value // reflect.Value of method Save of *grom.DB getted from *gorm.DB type
	methodWhere    reflect.Value // reflect.Value of method Where of *grom.DB getted from *gorm.DB type
	methodPreload  reflect.Value // reflect.Value of method Preload of *grom.DB getted from *gorm.DB type
	methodFind     reflect.Value // reflect.Value of method Find of *grom.DB getted from *gorm.DB type
	methodDelete   reflect.Value // reflect.Value of method Delete of *grom.DB getted from *gorm.DB type
	methodDebug    reflect.Value // reflect.Value of method Debug of *grom.DB getted from *gorm.DB type
}

// return pointer to gorm.DB
func (d *db) GormDB() interface{} {
	return d.gormDB.Interface()
}

// simulation for *gorm.DB method Limit
func (d *db) Limit(i int) types.IGormDB {
	results := d.methodLimit.
		Call([]reflect.Value{
			d.gormDB,
			reflect.ValueOf(i),
		})
	return d.clone(results[0])
}

// simulation for *gorm.DB method Offset
func (d *db) Offset(i int) types.IGormDB {
	results := d.methodOffset.
		Call([]reflect.Value{

			d.gormDB, reflect.ValueOf(i)})
	return d.clone(results[0])
}

// simulation for *gorm.DB method Count
func (d *db) Count(count *int64) types.IGormDB {
	results := d.methodCount.
		Call([]reflect.Value{

			d.gormDB, reflect.ValueOf(count)})
	return d.clone(results[0])
}

// simulation for *gorm.DB method Unscoped
func (d *db) Unscoped() types.IGormDB {
	results := d.methodUnscoped.
		Call([]reflect.Value{
			d.gormDB})
	return d.clone(results[0])
}

// simulation for *gorm.DB method Create
func (d *db) Create(value interface{}) types.IGormDB {
	results := d.methodCreate.
		Call([]reflect.Value{

			d.gormDB, reflect.ValueOf(value)})
	return d.clone(results[0])
}

// simulation for *gorm.DB method Save
func (d *db) Save(value interface{}) types.IGormDB {
	results := d.methodSave.
		Call([]reflect.Value{

			d.gormDB, reflect.ValueOf(value)})
	return d.clone(results[0])
}

// simulation for *gorm.DB method Where
func (d *db) Where(query interface{}, args ...interface{}) types.IGormDB {
	arguments := []reflect.Value{
		d.gormDB,
		reflect.ValueOf(query),
		reflect.ValueOf(args),
	}
	results := d.methodWhere.
		CallSlice(arguments)
	return d.clone(results[0])
}

// simulation for *gorm.DB method Preload
func (d *db) Preload(query string, args ...interface{}) types.IGormDB {
	arguments := []reflect.Value{
		d.gormDB,
		reflect.ValueOf(query),
		reflect.ValueOf(args),
	}
	results := d.methodPreload.
		CallSlice(arguments)
	return d.clone(results[0])
}

// simulation for *gorm.DB method Find
func (d *db) Find(dest interface{}, conds ...interface{}) types.IGormDB {
	arguments := []reflect.Value{
		d.gormDB,
		reflect.ValueOf(dest),
		reflect.ValueOf(conds),
	}
	results := d.methodFind.
		CallSlice(arguments)
	return d.clone(results[0])
}

// simulation for *gorm.DB method Delete
func (d *db) Delete(value interface{}, conds ...interface{}) types.IGormDB {
	arguments := []reflect.Value{
		d.gormDB,
		reflect.ValueOf(value),
		reflect.ValueOf(conds),
	}
	results := d.methodDelete.
		CallSlice(arguments)
	return d.clone(results[0])
}

// getter for field Error of *gorm.DB struct
func (d *db) Error() (err error) {
	val := d.gormDB.Elem().FieldByName("Error")
	if val.IsValid() && val.Interface() != nil {
		err = val.Interface().(error)
	}
	return
}

// getter for field RowsAffected of *gorm.DB struct
func (d *db) RowsAffected() (result int64) {
	val := d.gormDB.Elem().FieldByName("RowsAffected")
	if val.IsValid() && val.Interface() != nil {
		result = val.Interface().(int64)
	}
	return
}

// simulation for *gorm.DB method Debug
func (d *db) Debug() types.IGormDB {
	results := d.methodDebug.
		Call([]reflect.Value{
			d.gormDB})
	return d.clone(results[0])
}

// method clones db with calculated methods on parent db generated
func (d *db) clone(gormDB reflect.Value) types.IGormDB {
	return &db{
		gormDB:         gormDB,
		methodLimit:    d.methodLimit,
		methodOffset:   d.methodOffset,
		methodCount:    d.methodCount,
		methodUnscoped: d.methodUnscoped,
		methodCreate:   d.methodCreate,
		methodSave:     d.methodSave,
		methodWhere:    d.methodWhere,
		methodPreload:  d.methodPreload,
		methodFind:     d.methodFind,
		methodDelete:   d.methodDelete,
		methodDebug:    d.methodDebug,
	}
}

// Create types.IGormDB
func Db(data interface{}) (result types.IGormDB, err error) {
	dataType := reflect.TypeOf(data)
	if dataType.String() != "*gorm.DB" {
		err = errors.New("db is not *gorm.DB")
		return
	}

	methodLimit, okMethodLimit := dataType.MethodByName("Limit")
	if !okMethodLimit {
		err = errors.New("this struct have not method Limit")
		return
	}
	methodOffset, okMethodOffset := dataType.MethodByName("Offset")
	if !okMethodOffset {
		err = errors.New("this struct have not method Offset")
		return
	}
	methodCount, okMethodCount := dataType.MethodByName("Count")
	if !okMethodCount {
		err = errors.New("this struct have not method Count")
		return
	}
	methodUnscoped, okMethodUnscoped := dataType.MethodByName("Unscoped")
	if !okMethodUnscoped {
		err = errors.New("this struct have not method Unscoped")
		return
	}
	methodCreate, okMethodCreate := dataType.MethodByName("Create")
	if !okMethodCreate {
		err = errors.New("this struct have not method Create")
		return
	}
	methodSave, okMethodSave := dataType.MethodByName("Save")
	if !okMethodSave {
		err = errors.New("this struct have not method Save")
		return
	}
	methodWhere, okMethodWhere := dataType.MethodByName("Where")
	if !okMethodWhere {
		err = errors.New("this struct have not method Where")
		return
	}
	methodPreload, okMethodPreload := dataType.MethodByName("Preload")
	if !okMethodPreload {
		err = errors.New("this struct have not method Preload")
		return
	}
	methodFind, okMethodFind := dataType.MethodByName("Find")
	if !okMethodFind {
		err = errors.New("this struct have not method Find")
		return
	}
	methodDelete, okMethodDelete := dataType.MethodByName("Delete")
	if !okMethodDelete {
		err = errors.New("this struct have not method Delete")
		return
	}
	methodDebug, okMethodDebug := dataType.MethodByName("Debug")
	if !okMethodDebug {
		err = errors.New("this struct have not method Debug")
		return
	}

	return &db{
		gormDB:         reflect.ValueOf(data),
		methodLimit:    methodLimit.Func,
		methodOffset:   methodOffset.Func,
		methodCount:    methodCount.Func,
		methodUnscoped: methodUnscoped.Func,
		methodCreate:   methodCreate.Func,
		methodSave:     methodSave.Func,
		methodWhere:    methodWhere.Func,
		methodPreload:  methodPreload.Func,
		methodFind:     methodFind.Func,
		methodDelete:   methodDelete.Func,
		methodDebug:    methodDebug.Func,
	}, nil
}
