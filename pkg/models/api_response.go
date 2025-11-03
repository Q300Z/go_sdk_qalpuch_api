package models

// APIResponse represents the standard API response structure.
type APIResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    *interface{} `json:"data,omitempty"`
	Error   *interface{} `json:"error,omitempty"`
}

// APIErrorResponse represents the standard API error response structure.
type APIErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
