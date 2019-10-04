package httpext

import (
	"net/http"
)

// ResponseError - wrapper over error for http handlers
type ResponseError struct {
	err    error
	status int
}

func (re *ResponseError) Error() string {
	if re.err == nil {
		return ""
	}
	return re.err.Error()
}

// Status - http status code
func (re *ResponseError) Status() int {
	return re.status
}

// NoError - creates response error with status 200 OK
func NoError() *ResponseError {
	return &ResponseError{
		err:    nil,
		status: http.StatusOK,
	}
}

// ValidationError - creates response error with passed validation error and status code 400 BAD REQUEST
func ValidationError(err error) *ResponseError {
	return &ResponseError{
		err:    err,
		status: http.StatusBadRequest,
	}
}

// BadRequestError - creates response error with passed error and status code 400 BAD REQUEST
func BadRequestError(err error) *ResponseError {
	return &ResponseError{
		err:    err,
		status: http.StatusBadRequest,
	}
}

// InternalServerError - creates response error with internal server error and status code 500
func InternalServerError(err error) *ResponseError {
	return &ResponseError{
		err:    err,
		status: http.StatusInternalServerError,
	}
}

// OtherError - creates custom error response with custom error and status code
func OtherError(err error, code int) *ResponseError {
	return &ResponseError{
		err:    err,
		status: code,
	}
}
