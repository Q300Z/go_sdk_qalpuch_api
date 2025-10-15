package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/errors"
	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/models"
	"net/http"
)

// AuthClient handles authentication-related API requests.
type AuthClient struct {
	client *Client
}

// Login authenticates a user and returns a JWT.
func (c *AuthClient) Login(ctx context.Context, req models.LoginRequest) (*models.LoginResponse, error) {
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v1/login", c.client.BaseURL)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.client.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.NewAPIErrorFromResponse(resp)
	}

	var loginResp models.LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
		return nil, err
	}

	return &loginResp, nil
}

// Register creates a new user account.
func (c *AuthClient) Register(ctx context.Context, req models.RegisterRequest) (*models.RegisterResponse, error) {
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v1/register", c.client.BaseURL)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.client.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, errors.NewAPIErrorFromResponse(resp)
	}

	var registerResp models.RegisterResponse
	if err := json.NewDecoder(resp.Body).Decode(&registerResp); err != nil {
		return nil, err
	}

	return &registerResp, nil
}

// Logout invalidates the user's session.
func (c *AuthClient) Logout(ctx context.Context, req models.LogoutRequest) error {
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/v1/logout", c.client.BaseURL)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	httpReq.Header.Set("Content-Type", "application/json")
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

// ChangePassword allows a user to change their password.
func (c *AuthClient) ChangePassword(ctx context.Context, req models.ChangePasswordRequest) error {
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/v1/change-password", c.client.BaseURL)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	httpReq.Header.Set("Content-Type", "application/json")
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

func (c *AuthClient) RefreshToken(ctx context.Context, req models.RefreshTokenRequest) (*models.RefreshResponse, error) {
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v1/refresh", c.client.BaseURL)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.client.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.NewAPIErrorFromResponse(resp)
	}

	var refreshResp models.RefreshResponse
	if err := json.NewDecoder(resp.Body).Decode(&refreshResp); err != nil {
		return nil, err
	}

	return &refreshResp, nil
}
