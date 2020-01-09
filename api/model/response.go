package model

type Response struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func UnauthorizedResponse() Response {
	return Response{
		Success: false,
		Error:   "Unauthorized",
		Data:    nil,
	}
}

func ErrorResponse(err error) Response {
	return Response{
		Success: false,
		Error:   err.Error(),
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
