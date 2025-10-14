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

// WorkerClient handles worker-related API requests.
type WorkerClient struct {
	client *Client
}

// GetWorkers retrieves a list of all registered workers.
func (c *WorkerClient) GetWorkers(ctx context.Context) ([]models.Worker, error) {
	url := fmt.Sprintf("%s/v1/worker", c.client.BaseURL)
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

	var workersResp struct {
		Success bool            `json:"success"`
		Message string          `json:"message"`
		Data    []models.Worker `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&workersResp); err != nil {
		return nil, err
	}

	return workersResp.Data, nil
}

// GetWorker retrieves information about a specific worker.
func (c *WorkerClient) GetWorker(ctx context.Context, cuid string) (*models.Worker, error) {
	url := fmt.Sprintf("%s/v1/worker/%s", c.client.BaseURL, cuid)
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

	var workerResp struct {
		Success bool          `json:"success"`
		Message string        `json:"message"`
		Data    models.Worker `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&workerResp); err != nil {
		return nil, err
	}

	return &workerResp.Data, nil
}

// CreateWorker creates a new worker and returns a one-time registration token.
func (c *WorkerClient) CreateWorker(ctx context.Context, name string, capabilities []string) (*models.Worker, error) {
	requestBody := map[string]interface{}{"name": name, "capabilities": capabilities}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v1/worker", c.client.BaseURL)
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

	if resp.StatusCode != http.StatusCreated {
		return nil, errors.NewAPIErrorFromResponse(resp)
	}

	var workerResp struct {
		Success bool          `json:"success"`
		Message string        `json:"message"`
		Data    models.Worker `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&workerResp); err != nil {
		return nil, err
	}

	return &workerResp.Data, nil
}

// DeleteWorker deletes a worker by ID.
func (c *WorkerClient) DeleteWorker(ctx context.Context, cuid string) error {
	url := fmt.Sprintf("%s/v1/worker/%s", c.client.BaseURL, cuid)
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

// RegisterWorker is used by a worker to exchange its registration token for a persistent JWT.
func (c *WorkerClient) RegisterWorker(ctx context.Context, token string) (*models.AuthWorkerResponse, error) {
	requestBody := map[string]string{"token": token}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v1/worker/register", c.client.BaseURL)
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

	var authResp models.AuthWorkerResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return nil, err
	}

	return &authResp, nil
}

// RefreshAuth is used by a worker to refresh its JWT.
func (c *WorkerClient) RefreshAuth(ctx context.Context, refreshToken string) (*models.AuthWorkerResponse, error) {
	requestBody := map[string]string{"refreshToken": refreshToken}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v1/worker/refresh-auth", c.client.BaseURL)
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

	var authResp models.AuthWorkerResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return nil, err
	}

	return &authResp, nil
}
