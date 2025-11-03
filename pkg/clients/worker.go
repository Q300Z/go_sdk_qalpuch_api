package clients

import (
	"context"
	"fmt"

	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/models"
	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/services"
)

// WorkerClient implements services.WorkerService.
type WorkerClient struct {
	client *Client
}

// NewWorkerClient creates a new WorkerClient.
func NewWorkerClient(client *Client) services.WorkerService {
	return &WorkerClient{client: client}
}

// GetWorkers retrieves all workers.
func (c *WorkerClient) GetWorkers(ctx context.Context) ([]models.Worker, error) {
	workers := []models.Worker{}
	err := c.client.Get(ctx, "/worker", &workers)
	if err != nil {
		return nil, err
	}
	return workers, nil
}

// GetWorker retrieves a worker by ID.
func (c *WorkerClient) GetWorker(ctx context.Context, cuid string) (*models.Worker, error) {
	worker := &models.Worker{}
	err := c.client.Get(ctx, fmt.Sprintf("/worker/%s", cuid), worker)
	if err != nil {
		return nil, err
	}
	return worker, nil
}

// CreateWorker creates a new worker.
func (c *WorkerClient) CreateWorker(ctx context.Context, name string, capabilities []string) (*models.Worker, error) {
	req := models.CreateWorkerRequest{Name: name, Capabilities: capabilities}
	worker := &models.Worker{}
	err := c.client.Post(ctx, "/worker", req, worker)
	if err != nil {
		return nil, err
	}
	return worker, nil
}

// DeleteWorker deletes a worker.
func (c *WorkerClient) DeleteWorker(ctx context.Context, cuid string) error {
	return c.client.Delete(ctx, fmt.Sprintf("/worker/%s", cuid))
}

// RegisterWorker registers a worker.
func (c *WorkerClient) RegisterWorker(ctx context.Context, token string) (*models.AuthWorkerResponse, error) {
	req := models.RegisterWorkerRequest{Token: token}
	resp := &models.AuthWorkerResponse{}
	err := c.client.Post(ctx, "/worker/register", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RefreshAuth refreshes the worker's authentication token.
func (c *WorkerClient) RefreshAuth(ctx context.Context, refreshToken string) (*models.AuthWorkerResponse, error) {
	req := models.RefreshTokenRequest{RefreshToken: refreshToken}
	resp := &models.AuthWorkerResponse{}
	err := c.client.Post(ctx, "/worker/refresh-auth", req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
