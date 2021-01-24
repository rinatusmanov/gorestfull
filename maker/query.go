package maker

import (
	"errors"
	"github.com/rinatusmanov/gorestfull/types"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type query struct {
	keys     []string
	defaults []interface{}
	uri      *url.URL
	data     map[string]interface{}
}

func (u *query) URL() *url.URL {
	return u.uri
}

// Установка ключей по которым будет идти поиск в строке url
func (u *query) Keys(data ...string) types.IHttpDriverQuery {
	u.keys = data
	return u
}

// Установка значений по умолчанию
func (u *query) Defaults(data ...interface{}) types.IHttpDriverQuery {
	u.defaults = data
	return u
}

// возврат всех данных которые были запрошены
func (u *query) Data() map[string]interface{} {
	return u.data
}

// Парсинг данных из url в переменные которые переданы через ссылки
func (u *query) Parse(inData ...interface{}) (errSlice []error) {
	defer func() {
		u.keys = []string{}
		u.defaults = []interface{}{}
	}()
	if len(inData) != len(u.keys) {
		return []error{errors.New("keys and request inData have other length")}
	}
	var isDefault bool
	if len(u.defaults) == len(u.keys) {
		isDefault = true
	}
	for index, in := range inData {
		valPtr := reflect.ValueOf(in)
		val := valPtr
		if !val.IsValid() {
			return []error{errors.New("value is not valid")}
		}
		if valPtr.Kind() != reflect.Ptr {
			return []error{errors.New("interface is not pointer")}
		}
		val = val.Elem()
		if isDefault {
			defaultData := reflect.ValueOf(u.defaults[index])
			if defaultData.Type().Kind() != val.Type().Kind() {
				return []error{errors.New("type unreachable")}
			}
			val.Set(defaultData)
			u.data[u.keys[index]] = val.Interface()
		}
		if u.uri == nil {
			continue
		}
		strSl, okStr := u.uri.Query()[u.keys[index]]
		if okStr {
			_ = set(val, strSl)
			u.data[u.keys[index]] = val.Interface()
		}
	}
	return
}

// Присвоение значения через рефлексию
func set(val reflect.Value, stringSlice []string) (errSlice []error) {
	str := strings.Join(stringSlice, "")
	switch val.Kind() {
	case reflect.Bool:
		data, errParse := strconv.ParseBool(str)
		if errParse == nil {
			val.SetBool(data)
		}
		errSlice = append(errSlice, errParse)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		data, errParse := strconv.ParseInt(str, 0, 64)
		if errParse == nil {
			val.SetInt(data)
		}
		errSlice = append(errSlice, errParse)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		data, errParse := strconv.ParseUint(str, 0, 64)
		if errParse == nil {
			val.SetUint(data)
		}
		errSlice = append(errSlice, errParse)
	case reflect.Float32, reflect.Float64:
		data, errParse := strconv.ParseFloat(str, 64)
		if errParse == nil {
			val.SetFloat(data)
		}
		errSlice = append(errSlice, errParse)
	case reflect.String:
		val.SetString(str)
	case reflect.Ptr:
		return set(val.Elem(), stringSlice)
	case reflect.Array, reflect.Slice:
		newSlice := reflect.MakeSlice(val.Type(), len(stringSlice), len(stringSlice))
		newVal := val.Type().Elem()
		for i := 0; i < len(stringSlice); i++ {
			oneValueOfSlice := reflect.New(newVal)
			errOneValueSlice := set(oneValueOfSlice, strings.Split(stringSlice[i], ""))
			errSlice = append(errSlice, errOneValueSlice...)
			newSlice.Index(i).Set(oneValueOfSlice.Elem())
		}
		val.Set(newSlice)
	default:
		return []error{errors.New("undefined type")}
	}
	return
}

// Создает экземпляр query
func Query(uri *url.URL) types.IHttpDriverQuery {
	return &query{
		uri:  uri,
		data: make(map[string]interface{}),
	}
}
