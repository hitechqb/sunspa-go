package common

import "net/http"

type response struct {
	Code    int         `json:"code,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   error       `json:"error,omitempty,inline"`
	Message string      `json:"message,omitempty"`
}

func NewSuccessResponse(data interface{}) (int, interface{}) {
	return http.StatusOK, response{
		Code: 200,
		Data: data,
	}
}

func NewBadRequestResponse(err error, msg string) (int, interface{}) {
	if err == nil {
		err = NewError(ErrorInputBadRequest)
	}

	return http.StatusBadRequest, response{
		Code:    404,
		Error:   err,
		Message: msg,
	}
}

func NewUnauthorizedResponse(msg string) (int, interface{}) {
	return http.StatusUnauthorized, response{
		Code:    403,
		Error:   NewError(ErrorAuthorization),
		Message: msg,
	}
}

func NewInternalServerResponse(err error) (int, interface{}) {
	return http.StatusInternalServerError, response{
		Code:    500,
		Error:   err,
		Message: ErrorInternalServer.ToString(),
	}
}
