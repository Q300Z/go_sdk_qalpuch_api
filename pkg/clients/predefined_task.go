package clients

import (
	"context"
	"fmt"

	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/models"
)

// PredefinedTaskService defines the interface for predefined task operations.
type PredefinedTaskService interface {
	CreatePredefinedTask(ctx context.Context, req models.CreatePredefinedTaskRequest) (*models.PredefinedTask, error)
	GetPredefinedTasks(ctx context.Context) ([]models.PredefinedTask, error)
	GetPredefinedTaskByID(ctx context.Context, id string) (*models.PredefinedTask, error)
	UpdatePredefinedTask(ctx context.Context, id string, req models.UpdatePredefinedTaskRequest) (*models.PredefinedTask, error)
	DeletePredefinedTask(ctx context.Context, id string) error
}

// PredefinedTaskClient implements PredefinedTaskService.
type PredefinedTaskClient struct {
	client *Client
}

// NewPredefinedTaskClient creates a new PredefinedTaskClient.
func NewPredefinedTaskClient(client *Client) *PredefinedTaskClient {
	return &PredefinedTaskClient{client: client}
}

// CreatePredefinedTask creates a new predefined task.
func (c *PredefinedTaskClient) CreatePredefinedTask(ctx context.Context, req models.CreatePredefinedTaskRequest) (*models.PredefinedTask, error) {
	predefinedTask := &models.PredefinedTask{}
	err := c.client.Post(ctx, "/predefined-tasks", req, predefinedTask)
	if err != nil {
		return nil, err
	}
	return predefinedTask, nil
}

// GetPredefinedTasks retrieves all predefined tasks.
func (c *PredefinedTaskClient) GetPredefinedTasks(ctx context.Context) ([]models.PredefinedTask, error) {
	predefinedTasks := []models.PredefinedTask{}
	err := c.client.Get(ctx, "/predefined-tasks", &predefinedTasks)
	if err != nil {
		return nil, err
	}
	return predefinedTasks, nil
}

// GetPredefinedTaskByID retrieves a predefined task by its ID.
func (c *PredefinedTaskClient) GetPredefinedTaskByID(ctx context.Context, id string) (*models.PredefinedTask, error) {
	predefinedTask := &models.PredefinedTask{}
	err := c.client.Get(ctx, fmt.Sprintf("/predefined-tasks/%s", id), predefinedTask)
	if err != nil {
		return nil, err
	}
	return predefinedTask, nil
}

// UpdatePredefinedTask updates an existing predefined task.
func (c *PredefinedTaskClient) UpdatePredefinedTask(ctx context.Context, id string, req models.UpdatePredefinedTaskRequest) (*models.PredefinedTask, error) {
	predefinedTask := &models.PredefinedTask{}
	err := c.client.Put(ctx, fmt.Sprintf("/predefined-tasks/%s", id), req, predefinedTask)
	if err != nil {
		return nil, err
	}
	return predefinedTask, nil
}

// DeletePredefinedTask deletes a predefined task by its ID.
func (c *PredefinedTaskClient) DeletePredefinedTask(ctx context.Context, id string) error {
	return c.client.Delete(ctx, fmt.Sprintf("/predefined-tasks/%s", id))
}
