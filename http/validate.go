package http

import (
	"fmt"
	"net/http"
)

// ValidateMethod validates the HTTP method
func ValidateMethod(m string) error {
	switch m {
	case http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete,
		http.MethodConnect, http.MethodOptions, http.MethodTrace:
		return nil
	default:
		return fmt.Errorf("%s is not a HTTP method", m)
	}
}

// IsHTTPMethod indicates whether the specified string is a HTTP method
func IsHTTPMethod(m string) bool {
	return ValidateMethod(m) == nil
}
