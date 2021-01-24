package maker

import (
	"encoding/json"
	"fmt"
	"github.com/rinatusmanov/gorestfull/types"
	"reflect"
)

type response struct {
	Result     interface{} `json:"result,omitempty"`
	Count      int64       `json:"Count,omitempty"`
	AllCount   int64       `json:"all_count,omitempty"`
	Status     bool        `json:"Status,omitempty"`
	StatusCode uint64      `json:"status_code,omitempty"`
	ErrorStr   string      `json:"error,omitempty"`
	readWriter types.IReadWriter `json:"-"`
}

func (response *response) GetReadWriter() types.IReadWriter {
	return response.readWriter
}

func (response *response) SetReadWriter(readWriter types.IReadWriter) types.IResponse {
	response.readWriter = readWriter
	return response
}

func (response *response) GetErrorStr() string {
	return response.ErrorStr
}

func (response *response) SetErrorStr(errorStr string) types.IResponse {
	response.ErrorStr = errorStr
	return response
}

func (response *response) GetStatusCode() uint64 {
	return response.StatusCode
}

func (response *response) SetStatusCode(statusCode uint64) types.IResponse {
	response.StatusCode = statusCode
	return response
}

func (response *response) GetStatus() bool {
	return response.Status
}

func (response *response) SetStatus(status bool) types.IResponse {
	response.Status = status
	return response
}

func (response *response) GetAllCount() int64 {
	return response.AllCount
}

func (response *response) SetAllCount(allCount int64) types.IResponse {
	response.AllCount = allCount
	return response
}

func (response *response) GetCount() int64 {
	return response.Count
}

func (response *response) SetCount(count int64) types.IResponse {
	response.Count = count
	return response
}

func (response *response) GetResult() interface{} {
	return response.Result
}

func (response *response) SetResult(result interface{}) types.IResponse {
	response.Result = result
	return response
}

func (response *response) Write() (count int, err error) {
	byteSl, errMarshal := json.Marshal(response)
	if errMarshal != nil {
		count, err = response.readWriter.Write([]byte(`{"Status":false,"status_code":399,"error":"` + errMarshal.Error() + `"}`))
	} else {
		count, err = response.readWriter.Write(byteSl)
	}
	return
}

func (response *response) Parse(result reflect.Value, tx types.IGormDB) types.IResponse {
	var countOfQuery int64
	if tx.Error() != nil {
		response.Error(tx.Error)
	} else {
		response.Ok(result)
		response.Count = tx.RowsAffected()
		tx.Count(&countOfQuery)
		response.AllCount = countOfQuery
	}
	return response
}

func (response *response) Ok(data reflect.Value) types.IResponse {
	response.Status = true
	response.StatusCode = 200
	response.Result = data.Interface()
	return response
}

func (response *response) Error(data interface{}) types.IResponse {
	if data == nil {
		return response
	}
	response.Status = false
	response.StatusCode = 400
	response.ErrorStr = fmt.Sprint(data)
	return response
}

func NewResponse(readWriter types.IReadWriter) types.IResponse {
	return &response{readWriter: readWriter}
}
