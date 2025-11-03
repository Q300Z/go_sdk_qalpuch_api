package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/errors"
	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/models"
	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/services"
)

// Client manages communication with the Qalpuch API.
type Client struct {
	HTTPClient *http.Client
	BaseURL    string
	Token      string // JWT token for authentication

	Auth            services.AuthService
	Users           services.UserService
	Files           services.FileService
	Tasks           services.TaskService
	Workers         services.WorkerService
	PredefinedTasks services.PredefinedTaskService
}

// NewClient creates a new API client.
func NewClient(baseURL, token string) *Client {
	c := &Client{
		BaseURL: baseURL,
		Token:   token,
		HTTPClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}

	// Initialize sub-clients
	c.Auth = NewAuthClient(c)
	c.Users = NewUserClient(c)
	c.Files = NewFileClient(c)
	c.Tasks = NewTaskClient(c)
	c.Workers = NewWorkerClient(c)
	c.PredefinedTasks = NewPredefinedTaskClient(c)

	return c
}

// Request performs an HTTP request to the API.
func (c *Client) Request(ctx context.Context, method, path string, reqBody, respBody interface{}) error {
	var body io.Reader
	if reqBody != nil {
		reqBytes, err := json.Marshal(reqBody)
		if err != nil {
			return fmt.Errorf("failed to marshal request body for %s %s: %w", method, path, err)
		}
		body = bytes.NewBuffer(reqBytes)
	}

	url := fmt.Sprintf("%s%s", c.BaseURL, path)
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return fmt.Errorf("failed to create request for %s %s: %w", method, url, err)
	}

	req.Header.Set("Content-Type", "application/json")
	if c.Token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to perform %s request to %s: %w", method, url, err)
	}
	defer resp.Body.Close()

	var apiResponse models.APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		// If we can't decode into APIResponse, try to decode into APIErrorResponse for error handling
		var apiErrResp models.APIErrorResponse
		if err := json.NewDecoder(bytes.NewBufferString(resp.Status)).Decode(&apiErrResp); err == nil {
			return &errors.APIError{StatusCode: resp.StatusCode, Message: fmt.Sprintf("API error for %s %s: %s", method, url, apiErrResp.Message)}
		}
		return fmt.Errorf("failed to decode API response for %s %s: %w", method, url, err)
	}

	if !apiResponse.Success {
		var errMsg string
		if apiResponse.Message != "" {
			errMsg = apiResponse.Message
		} else if apiResponse.Error != nil {
			// Attempt to extract error message from the Error field if it's a string
			if errStr, ok := (*apiResponse.Error).(string); ok {
				errMsg = errStr
			} else {
				// Fallback if Error is not a string
				errMsg = fmt.Sprintf("API error with status %d for %s %s", resp.StatusCode, method, url)
			}
		} else {
			errMsg = fmt.Sprintf("API error with status %d for %s %s", resp.StatusCode, method, url)
		}
		return &errors.APIError{StatusCode: resp.StatusCode, Message: errMsg}
	}

	if respBody != nil && apiResponse.Data != nil {
		dataBytes, err := json.Marshal(apiResponse.Data)
		if err != nil {
			return fmt.Errorf("failed to marshal API response data for %s %s: %w", method, url, err)
		}
		if err := json.Unmarshal(dataBytes, respBody); err != nil {
			return fmt.Errorf("failed to unmarshal API response data for %s %s: %w", method, url, err)
		}
	}

	return nil
}

// Get performs a GET request.
func (c *Client) Get(ctx context.Context, path string, respBody interface{}) error {
	return c.Request(ctx, http.MethodGet, path, nil, respBody)
}

// Post performs a POST request.
func (c *Client) Post(ctx context.Context, path string, reqBody, respBody interface{}) error {
	return c.Request(ctx, http.MethodPost, path, reqBody, respBody)
}

// Put performs a PUT request.
func (c *Client) Put(ctx context.Context, path string, reqBody, respBody interface{}) error {
	return c.Request(ctx, http.MethodPut, path, reqBody, respBody)
}

// Patch performs a PATCH request.
func (c *Client) Patch(ctx context.Context, path string, reqBody, respBody interface{}) error {
	return c.Request(ctx, http.MethodPatch, path, reqBody, respBody)
}

// Delete performs a DELETE request.
func (c *Client) Delete(ctx context.Context, path string) error {
	return c.Request(ctx, http.MethodDelete, path, nil, nil)
}
