package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go_sdk_qalpuch_api/pkg/errors"
	"go_sdk_qalpuch_api/pkg/models"
	"net/http"
)

// UserClient handles user-related API requests.
type UserClient struct {
	client *Client
}

// GetUsers retrieves a list of all users.
func (c *UserClient) GetUsers(ctx context.Context) ([]models.User, error) {
	url := fmt.Sprintf("%s/v1/users", c.client.BaseURL)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.client.Token))

	resp, err := c.client.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.NewAPIErrorFromResponse(resp)
	}

	var usersResp struct {
		Success bool          `json:"success"`
		Message string        `json:"message"`
		Data    []models.User `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&usersResp); err != nil {
		return nil, err
	}

	return usersResp.Data, nil
}

// GetUser retrieves a single user by ID.
func (c *UserClient) GetUser(ctx context.Context, id int) (*models.User, error) {
	url := fmt.Sprintf("%s/v1/users/%d", c.client.BaseURL, id)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.client.Token))

	resp, err := c.client.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.NewAPIErrorFromResponse(resp)
	}

	var userResp struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    models.User `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&userResp); err != nil {
		return nil, err
	}

	return &userResp.Data, nil
}

// UpdateUser updates a user's information.
func (c *UserClient) UpdateUser(ctx context.Context, id int, user models.User) (*models.User, error) {
	jsonBody, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v1/users/%d", c.client.BaseURL, id)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.client.Token))

	resp, err := c.client.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.NewAPIErrorFromResponse(resp)
	}

	var userResp struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    models.User `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&userResp); err != nil {
		return nil, err
	}

	return &userResp.Data, nil
}

// DeleteUser deletes a user by ID.
func (c *UserClient) DeleteUser(ctx context.Context, id int) error {
	url := fmt.Sprintf("%s/v1/users/%d", c.client.BaseURL, id)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.client.Token))

	resp, err := c.client.HTTPClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.NewAPIErrorFromResponse(resp)
	}

	return nil
}

// DeleteCurrentUser deletes the currently authenticated user.
func (c *UserClient) DeleteCurrentUser(ctx context.Context) error {
	url := fmt.Sprintf("%s/v1/users/me", c.client.BaseURL)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.client.Token))

	resp, err := c.client.HTTPClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.NewAPIErrorFromResponse(resp)
	}

	return nil
}

// CreateUser creates a new user (admin only).
func (c *UserClient) CreateUser(ctx context.Context, req models.CreateUserRequest) (*models.User, error) {
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v1/users", c.client.BaseURL)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.client.Token))

	resp, err := c.client.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.NewAPIErrorFromResponse(resp)
	}

	var userResp struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    models.User `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&userResp); err != nil {
		return nil, err
	}

	return &userResp.Data, nil
}

// SearchUsers searches for users by a query string (admin only).
func (c *UserClient) SearchUsers(ctx context.Context, query string) ([]models.User, error) {
	url := fmt.Sprintf("%s/v1/users/search?q=%s", c.client.BaseURL, query)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.client.Token))

	resp, err := c.client.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.NewAPIErrorFromResponse(resp)
	}

	var usersResp struct {
		Success bool          `json:"success"`
		Message string        `json:"message"`
		Data    []models.User `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&usersResp); err != nil {
		return nil, err
	}

	return usersResp.Data, nil
}
