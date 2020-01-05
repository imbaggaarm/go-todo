package model

type Response struct {
	Success bool
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
