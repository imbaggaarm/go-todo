package errformatter

import "errors"

type ErrorType string

const (
	ErrorDatabaseConnection ErrorType = "connection error"
	ErrorUserExisted        ErrorType = "user existed"
	ErrorWrongPassword      ErrorType = "wrong password"
	ErrorRecordNotFound     ErrorType = "record not found"
	ErrorUnauthorized       ErrorType = "unauthorized"
)

func NewError(error ErrorType) error {
	return errors.New(string(error))
}
