package clients

import (
	"net/http"
	"time"

	"go_sdk_qalpuch_api/pkg/services"
)

const (
	DefaultTimeout = 10 * time.Second
)

// Client manages communication with the QAlpuch API.
type Client struct {
	HTTPClient *http.Client
	BaseURL    string
	Token      string

	Auth    services.AuthService
	Users   services.UserService
	Files   services.FileService
	Tasks   services.TaskService
	Workers services.WorkerService
}

// NewClient creates a new API client.
func NewClient(baseURL string, token string) *Client {
	return &Client{
		HTTPClient: &http.Client{Timeout: DefaultTimeout},
		BaseURL:    baseURL,
		Token:      token,
	}
}

// NewUserClient creates a new client for user-facing operations.
func NewUserClient(baseURL string, token string) *Client {
	c := NewClient(baseURL, token)
	c.Auth = &AuthClient{client: c}
	c.Users = &UserClient{client: c}
	c.Files = &FileClient{client: c}
	c.Tasks = &TaskClient{client: c}
	c.Workers = &WorkerClient{client: c}
	return c
}

// NewWorkerClient creates a new client for worker-specific operations.
func NewWorkerClient(baseURL string, token string) *Client {
	c := NewClient(baseURL, token)
	c.Auth = &AuthClient{client: c}
	c.Tasks = &TaskClient{client: c}
	c.Workers = &WorkerClient{client: c}
	return c
}