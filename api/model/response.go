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

