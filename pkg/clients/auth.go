package clients

import (
	"context"

	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/models"
	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/services"
)

// AuthClient implements services.AuthService.
type AuthClient struct {
	client *Client // This Client is the one from client.go
}

// NewAuthClient creates a new AuthClient.
func NewAuthClient(client *Client) services.AuthService {
	return &AuthClient{client: client}
}

// Login authenticates a user.
func (c *AuthClient) Login(ctx context.Context, req models.LoginRequest) (*models.LoginResponse, error) {
	resp := &models.LoginResponse{}
	err := c.client.Post(ctx, "/login", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Register registers a new user.
func (c *AuthClient) Register(ctx context.Context, req models.RegisterRequest) (*models.LoginResponse, error) {
	resp := &models.LoginResponse{}
	err := c.client.Post(ctx, "/register", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Logout logs out the current user.
func (c *AuthClient) Logout(ctx context.Context, req models.LogoutRequest) error {
	return c.client.Post(ctx, "/logout", req, nil)
}

// ChangePassword changes the password of the authenticated user.
func (c *AuthClient) ChangePassword(ctx context.Context, req models.ChangePasswordRequest) error {
	return c.client.Post(ctx, "/change-password", req, nil)
}

// RefreshToken refreshes the access token using a refresh token.
func (c *AuthClient) RefreshToken(ctx context.Context, req models.RefreshTokenRequest) (*models.RefreshResponse, error) {
	resp := &models.RefreshResponse{}
	err := c.client.Post(ctx, "/refresh", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
