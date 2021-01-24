package types

import "net/url"

// IHttpDriverQuery упрощает работу со строкой запроса
type IHttpDriverQuery interface {
	// Keys устанавливает ключи по которым будет парсится строка запроса
	Keys(data ...string) IHttpDriverQuery
	// Defaults устанавливает значения по умолчанию количество значений по умолчанию должно соответствовать количеству ключей
	Defaults(data ...interface{}) IHttpDriverQuery
	// Parse парсит в переменные значения из строки запроса
	Parse(inData ...interface{}) (errSlice []error)
	// Data возвращает мапу где храняться все запрошенные до этого момента переменные
	Data() map[string]interface{}
	// URL возвращает url по которому идет запрос
	URL() *url.URL
}
