package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/errors"
	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/models"
	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/services"
	"github.com/go-playground/validator/v10"
	"io"
	"mime/multipart"
	"net/http"
	"regexp"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	if err := validate.RegisterValidation("regexp", func(fl validator.FieldLevel) bool {
		if fl.Field().String() == "" {
			return true // omitempty handled by other tags
		}
		_, err := regexp.MatchString(fl.Param(), fl.Field().String())
		return err == nil
	}); err != nil {
		panic(err)
	}
}

// TaskClient handles task-related API requests.
type TaskClient struct {
	client *Client
}

// Build starts the construction of a new task using a fluent builder pattern.
func (c *TaskClient) Build(fileID string) services.TaskBuilder {
	return &TaskBuilder{
		client: c,
		fileID: fileID,
	}
}

// GetUserTasks retrieves a list of all tasks created by the authenticated user.
func (c *TaskClient) GetUserTasks(ctx context.Context) ([]models.Task, error) {
	url := fmt.Sprintf("%s/v1/tasks", c.client.BaseURL)
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

	var tasksResp struct {
		Success bool          `json:"success"`
		Message string        `json:"message"`
		Data    []models.Task `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tasksResp); err != nil {
		return nil, err
	}

	return tasksResp.Data, nil
}

// CreateTask creates a new task from an existing file.
func (c *TaskClient) CreateTask(ctx context.Context, req models.CreateTaskRequest) (*models.Task, error) {
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v1/tasks", c.client.BaseURL)
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

	var taskResp struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    models.Task `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&taskResp); err != nil {
		return nil, err
	}

	return &taskResp.Data, nil
}

// DeleteTask deletes a task.
func (c *TaskClient) DeleteTask(ctx context.Context, cuid string) error {
	url := fmt.Sprintf("%s/v1/tasks/%s", c.client.BaseURL, cuid)
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

// GetPendingTask assigns a pending task to an authenticated worker.
func (c *TaskClient) GetPendingTask(ctx context.Context) (*models.Task, error) {
	url := fmt.Sprintf("%s/v1/tasks/pending", c.client.BaseURL)
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

	var taskResp struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    models.Task `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&taskResp); err != nil {
		return nil, err
	}

	return &taskResp.Data, nil
}

// UpdateTaskStatus allows a worker to update the status of a task.
func (c *TaskClient) UpdateTaskStatus(ctx context.Context, cuid string, req models.UpdateTaskStatusRequest) error {
	jsonBody, err := json.Marshal(req)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/v1/tasks/%s", c.client.BaseURL, cuid)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPatch, url, bytes.NewBuffer(jsonBody))
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

// UploadTaskResult allows a worker to upload the result of a completed task.
func (c *TaskClient) UploadTaskResult(ctx context.Context, cuid string, file []byte) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", "result.txt") // Assuming a generic filename for the result
	if err != nil {
		return err
	}

	_, err = io.Copy(part, bytes.NewReader(file))
	if err != nil {
		return err
	}

	writer.Close()

	url := fmt.Sprintf("%s/v1/tasks/%s/result", c.client.BaseURL, cuid)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return err
	}

	httpReq.Header.Set("Content-Type", writer.FormDataContentType())
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
