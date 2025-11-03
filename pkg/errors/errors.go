package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// APIError represents a detailed error returned by the QAlpuch API.
type APIError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Err        error  // Wrapped error
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error (status %d): %s", e.StatusCode, e.Message)
}

func (e *APIError) Unwrap() error {
	return e.Err
}

// Sentinel errors for common API issues.
var (
	ErrBadRequest     = errors.New("bad request")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidden      = errors.New("forbidden")
	ErrNotFound       = errors.New("not found")
	ErrInternalServer = errors.New("internal server error")
	ErrUnknown        = errors.New("an unknown API error occurred")
)

// NewAPIErrorFromResponse creates a new APIError from an HTTP response.
// It attempts to parse the error details from the response body.
