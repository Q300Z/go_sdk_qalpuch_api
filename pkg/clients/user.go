package clients

import (
	"context"
	"fmt"

	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/models"
	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/services"
)

// UserClient implements services.UserService.
type UserClient struct {
	client *Client
}

// NewUserClient creates a new UserClient.
func NewUserClient(client *Client) services.UserService {
	return &UserClient{client: client}
}

// GetUsers retrieves all users.
func (c *UserClient) GetUsers(ctx context.Context) ([]models.User, error) {
	users := []models.User{}
	err := c.client.Get(ctx, "/users", &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetUser retrieves a user by ID.
func (c *UserClient) GetUser(ctx context.Context, id int) (*models.User, error) {
	user := &models.User{}
	err := c.client.Get(ctx, fmt.Sprintf("/users/%d", id), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser updates a user.
func (c *UserClient) UpdateUser(ctx context.Context, id int, req models.UpdateUserRequest) (*models.User, error) {
	user := &models.User{}
	err := c.client.Put(ctx, fmt.Sprintf("/users/%d", id), req, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteUser deletes a user by ID.
func (c *UserClient) DeleteUser(ctx context.Context, id int) error {
	return c.client.Delete(ctx, fmt.Sprintf("/users/%d", id))
}

// DeleteCurrentUser deletes the authenticated user.
func (c *UserClient) DeleteCurrentUser(ctx context.Context) error {
	return c.client.Delete(ctx, "/users/me")
}

// CreateUser creates a new user.
func (c *UserClient) CreateUser(ctx context.Context, req models.CreateUserRequest) (*models.User, error) {
	user := &models.User{}
	err := c.client.Post(ctx, "/users", req, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// SearchUsers searches for users by query.
func (c *UserClient) SearchUsers(ctx context.Context, query string) ([]models.User, error) {
	users := []models.User{}
	err := c.client.Get(ctx, fmt.Sprintf("/users/search?q=%s", query), &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
