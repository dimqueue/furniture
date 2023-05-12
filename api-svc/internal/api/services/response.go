package services

import "net/http"

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

func NewSuccessResponseWithData(data interface{}) Response {
	return Response{
		Success: true,
		Data:    data,
	}
}

func NewErrorResponse(err error) Response {
	return Response{
		Success: false,
		Error:   err.Error(),
	}
}

func NewStringErrorResponse(errStr string) Response {
	return Response{
		Success: false,
		Error:   errStr,
	}
}

func NewDefaultResponse(httpStatusCode int) Response {
	return Response{
		Success: false,
		Error:   http.StatusText(httpStatusCode),
	}
}
