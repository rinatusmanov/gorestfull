package types

// IReadWriter the main task of the interface is to maintain input and output
// IReadWriter основная задача интерфейса обслуживание ввода и вывода
type IReadWriter interface {
	// Method возвращает http метод в параметре method
	// Method returns http method at method field
	Method() (method string)
	// GetBody возвращает тело запроса в параметре body
	// GetBody returns body of request at body field
	GetBody() (body []byte)
	// Write Пишет slice байтов(аргумент body) в ответ
	// Write method write to response body
	Write(body []byte) (count int, err error)
	// Query возвращает структуру IHttpDriverQuery, которая упрощает обработку строки запроса
	// Query returns struct IHttpDriverQuery
	Query() (query IHttpDriverQuery)
	// Response Возвращает экземпляр IResponse
	// Response returns IResponse
	Response() IResponse
}
