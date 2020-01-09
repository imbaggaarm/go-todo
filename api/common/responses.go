package common

import . "go-todo/api/util/errformatter"

type Response struct {
	Success bool        `json:"success"`
	Error   ErrorType   `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func UnauthorizedResponse() Response {
	return Response{
		Success: false,
		Error:   ErrorUnauthorized,
		Data:    nil,
	}
}

func ErrorResponse(err error) Response {
	return Response{
		Success: false,
		Error:   ErrorType(err.Error()),
		Data:    nil,
	}
}

func SuccessResponse(data interface{}) Response {
	return Response{
		Success: true,
		Error:   "",
		Data:    data,
	}
}
