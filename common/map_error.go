package common

import "errors"

type TypeError string

const (
	ErrorInputBadRequest TypeError = "Input Bad Request"
	ErrorAuthorization   TypeError = "Authorization Token is required"
	ErrorInternalServer  TypeError = "Internal Server"
)

func (s TypeError) ToString() string {
	return string(s)
}

func NewError(typeErr TypeError) error {
	return errors.New(typeErr.ToString())
}
