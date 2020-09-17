package reswrapper

import (
	"net/http"
)

type ResponseWrapper struct {
	Status     string      `json:"status"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"error,omitempty"`
	StatusCode int         `json:"-"`
}

func ErrorInputValidation() (errResponse ResponseWrapper) {
	errResponse.StatusCode = http.StatusBadRequest
	errResponse.Status = "FAIL"
	return
}

func InternalServerError(err error) (errResponse ResponseWrapper) {
	errResponse.StatusCode = http.StatusInternalServerError
	errResponse.Status = "FAIL"
	errResponse.Error = err.Error()
	return
}

func OK() (response ResponseWrapper) {
	response.Status = "OK"
	response.StatusCode = http.StatusOK
	return
}

func ErrorUnauthorized() (errResponse ResponseWrapper) {
	errResponse.StatusCode = http.StatusUnauthorized
	errResponse.Status = "FAIL"
	errResponse.Error = "Access Denied!"
	return
}
