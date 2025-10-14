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
func NewAPIErrorFromResponse(resp *http.Response) error {
	apiErr := &APIError{
		StatusCode: resp.StatusCode,
	}

	// Try to decode the standard error response from the API.
	body, err := io.ReadAll(resp.Body)
	if err == nil && len(body) > 0 {
		var errResp struct {
			Message string `json:"message"`
		}
		if json.Unmarshal(body, &errResp) == nil && errResp.Message != "" {
			apiErr.Message = errResp.Message
		} else {
			// If decoding fails, use the raw body as the message.
			apiErr.Message = string(body)
		}
	} else {
		// Fallback to the status text if reading the body fails.
		apiErr.Message = http.StatusText(resp.StatusCode)
	}

	switch resp.StatusCode {
	case http.StatusBadRequest:
		apiErr.Err = ErrBadRequest
	case http.StatusUnauthorized:
		apiErr.Err = ErrUnauthorized
	case http.StatusForbidden:
		apiErr.Err = ErrForbidden
	case http.StatusNotFound:
		apiErr.Err = ErrNotFound
	case http.StatusInternalServerError:
		apiErr.Err = ErrInternalServer
	default:
		apiErr.Err = ErrUnknown
	}

	return apiErr
}
