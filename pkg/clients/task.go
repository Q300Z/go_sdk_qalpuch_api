package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/errors"
	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/models"
	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/services"
)

// TaskClient implements services.TaskService.
type TaskClient struct {
	client *Client
}

// NewTaskClient creates a new TaskClient.
func NewTaskClient(client *Client) services.TaskService {
	return &TaskClient{client: client}
}

// GetUserTasks retrieves tasks for the authenticated user.
func (c *TaskClient) GetUserTasks(ctx context.Context) ([]models.Task, error) {
	tasks := []models.Task{}
	err := c.client.Get(ctx, "/tasks", &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// CreateTask creates a new task.
func (c *TaskClient) CreateTask(ctx context.Context, req models.CreateTaskRequest) (*models.Task, error) {
	task := &models.Task{}
	err := c.client.Post(ctx, "/tasks", req, task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

// DeleteTask deletes a task.
func (c *TaskClient) DeleteTask(ctx context.Context, cuid string) error {
	return c.client.Delete(ctx, fmt.Sprintf("/tasks/%s", cuid))
}

// GetPendingTask retrieves a pending task for a worker.
func (c *TaskClient) GetPendingTask(ctx context.Context) (*models.Task, error) {
	task := &models.Task{}
	err := c.client.Get(ctx, "/tasks/pending", task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

// UpdateTaskStatus updates the status of a task.
func (c *TaskClient) UpdateTaskStatus(ctx context.Context, cuid string, req models.UpdateTaskStatusRequest) error {
	return c.client.Patch(ctx, fmt.Sprintf("/tasks/%s", cuid), req, nil)
}

// UploadTaskResult uploads the result of a task.
func (c *TaskClient) UploadTaskResult(ctx context.Context, cuid string, filename string, file []byte) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return fmt.Errorf("failed to create form file: %w", err)
	}
	if _, err := part.Write(file); err != nil {
		return fmt.Errorf("failed to write file content: %w", err)
	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("failed to close multipart writer: %w", err)
	}

	url := fmt.Sprintf("%s/tasks/%s/result", c.client.BaseURL, cuid)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	if c.client.Token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.client.Token))
	}

	resp, err := c.client.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to perform request: %w", err)
	}
	defer resp.Body.Close()

	var apiResponse models.APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return fmt.Errorf("failed to decode API response: %w", err)
	}

	if !apiResponse.Success {
		return &errors.APIError{StatusCode: resp.StatusCode, Message: fmt.Sprintf("API error: %s", *apiResponse.Error)}
	}

	return nil
}

// Build creates a new TaskBuilder.
func (c *TaskClient) Build(fileID string) services.TaskBuilder {
	return NewTaskBuilder(c, fileID)
}
