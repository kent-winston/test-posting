package tools

import "net/http"

type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

func NewCustomError(code int, message string) error {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

func APIErrorResponse(err error) (int, string) {
	customErr, ok := err.(*CustomError)
	if ok {
		return customErr.Code, customErr.Message
	} else {
		return http.StatusInternalServerError, err.Error()
	}
}
