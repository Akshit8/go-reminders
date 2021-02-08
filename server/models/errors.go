package models

import "fmt"

// HTTPError is http error returned to the client
type HTTPError struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

// Error returns message field of HTTPError struct
func (e HTTPError) Error() string {
	return e.Message
}

// FormatValidationError represents the error returned in case the request body has
//  a wrong format which the server cannot work with
type FormatValidationError struct {
	Message string
}

// Error returns message field of FormatValidationError struct
func (e FormatValidationError) Error() string {
	return e.Message
}

// DataValidationError represents the error returned when the format of request
// is valid but the data is invalid
type DataValidationError struct {
	Message string
}

// Error returns message field of DataValidationError struct
func (e DataValidationError) Error() string {
	return e.Message
}

// InvalidJSONError represents the error returned when request body contains invalid JSON
type InvalidJSONError struct {
	Message string
}

// Error returns message field of InvalidJSONError struct
func (e InvalidJSONError) Error() string {
	return e.Message
}

// NotFoundError represents the error returned in case a resource or route is not found
type NotFoundError struct {
	Message string
}

// Error returns message field of NotFoundError struct
func (e NotFoundError) Error() string {
	if e.Message == "" {
		return "resource not found"
	}
	return e.Message
}

// WrapError wraps a plain error into a custom error
func WrapError(customErr string, originalErr error) error {
	err := fmt.Errorf("%s: %v", customErr, originalErr)
	return err
}
