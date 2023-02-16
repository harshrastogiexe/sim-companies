package models

import (
	"encoding/json"
)

type ApiError struct {
	// HTTP status codes as registered with IANA.
	//
	// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
	Code int

	Message string
}

func NewApiError(code int, message string) ApiError {
	return ApiError{code, message}
}

// convert ApiError instance to
//
//	{
//		"code": 500
//		"message": "Internal Server Error"
//	}
func (ae ApiError) JSON() ([]byte, error) {
	return json.Marshal(ae)
}
