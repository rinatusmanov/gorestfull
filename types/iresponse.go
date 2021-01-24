package types

import "reflect"


// IResponse стандартный ответ
type IResponse interface {
	// GetReadWriter получает значение поля ReadWriter
	GetReadWriter() IReadWriter
	// SetReadWriter устанавливает значение поля ReadWriter
	SetReadWriter(readWriter IReadWriter) IResponse
	// GetErrorStr получает значение поля ErrorStr
	GetErrorStr() string
	// SetErrorStr устанавливает значение поля ErrorStr
	SetErrorStr(errorStr string) IResponse
	// GetStatusCode получает значение поля StatusCode
	GetStatusCode() uint64
	// SetStatusCode устанавливает значение поля StatusCode
	SetStatusCode(statusCode uint64) IResponse
	// GetStatus получает значение поля Status
	GetStatus() bool
	// SetStatus устанавливает значение поля Status
	SetStatus(status bool) IResponse
	// GetAllCount получает значение поля AllCount
	GetAllCount() int64
	// SetAllCount устанавливает значение поля AllCount
	SetAllCount(allCount int64) IResponse
	// GetCount получает значение поля Count
	GetCount() int64
	// SetCount устанавливает значение поля Count
	SetCount(count int64) IResponse
	// GetResult получает значение поля Result
	GetResult() interface{}
	// SetResult устанавливает значение поля Result
	SetResult(result interface{}) IResponse
	// Write пишет свое тело в json в ответ
	Write() (count int, err error)
	// Parse формирует тело запроса
	Parse(result reflect.Value, tx IGormDB) IResponse
	// Ok формирует стандартный ответ
	Ok(data reflect.Value) IResponse
	// Error формирует ответ с ошибкой
	Error(data interface{}) IResponse
}
