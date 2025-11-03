package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/errors"
	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/models"
	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/services"
)

// FileClient implements services.FileService.
type FileClient struct {
	client *Client
}

// NewFileClient creates a new FileClient.
func NewFileClient(client *Client) services.FileService {
	return &FileClient{client: client}
}

// UploadFile uploads a file.
func (c *FileClient) UploadFile(ctx context.Context, name string, file []byte) (*models.File, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", name)
	if err != nil {
		return nil, fmt.Errorf("failed to create form file: %w", err)
	}
	if _, err := part.Write(file); err != nil {
		return nil, fmt.Errorf("failed to write file content: %w", err)
	}

	// Add other fields if necessary, e.g., "name"
	if err := writer.WriteField("name", name); err != nil {
		return nil, fmt.Errorf("failed to write name field: %w", err)
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close multipart writer: %w", err)
	}

	url := fmt.Sprintf("%s/files/upload", c.client.BaseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create upload request to %s: %w", url, err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	if c.client.Token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.client.Token))
	}

	resp, err := c.client.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform upload request to %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		var apiErr models.APIErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
			return nil, fmt.Errorf("API error with status %d for upload to %s: %s", resp.StatusCode, url, resp.Status)
		}
		return nil, &errors.APIError{StatusCode: resp.StatusCode, Message: fmt.Sprintf("API error for upload to %s: %s", url, apiErr.Message)}
	}

	var apiResponse models.APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode API response for upload to %s: %w", url, err)
	}

	fileResp := &models.File{}
	// The Data field in APIResponse is interface{}, so we need to re-marshal and unmarshal
	// to get the concrete type.
	dataBytes, err := json.Marshal(apiResponse.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal API response data for upload to %s: %w", url, err)
	}
	if err := json.Unmarshal(dataBytes, fileResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal file response data for upload to %s: %w", url, err)
	}

	return fileResp, nil
}

// GetFileMetadata retrieves file metadata.
func (c *FileClient) GetFileMetadata(ctx context.Context, cuid string) (*models.File, error) {
	var resp struct {
		Data models.File `json:"data"`
	}
	err := c.client.Get(ctx, fmt.Sprintf("/files/%s", cuid), &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}

// DownloadFile downloads a file.
func (c *FileClient) DownloadFile(ctx context.Context, cuid string) ([]byte, error) {
	url := fmt.Sprintf("%s/files/%s/download", c.client.BaseURL, cuid)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create download request to %s: %w", url, err)
	}

	if c.client.Token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.client.Token))
	}

	resp, err := c.client.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform download request to %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		var apiErr models.APIErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
			return nil, fmt.Errorf("API error with status %d for download from %s: %s", resp.StatusCode, url, resp.Status)
		}
		return nil, &errors.APIError{StatusCode: resp.StatusCode, Message: fmt.Sprintf("API error for download from %s: %s", url, apiErr.Message)}
	}

	return io.ReadAll(resp.Body)
}

// ListUserFiles lists files for the authenticated user.
func (c *FileClient) ListUserFiles(ctx context.Context) ([]models.File, error) {
	var resp struct {
		Data []models.File `json:"data"`
	}
	err := c.client.Get(ctx, "/files", &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// DeleteFile deletes a file.
func (c *FileClient) DeleteFile(ctx context.Context, cuid string) error {
	return c.client.Delete(ctx, fmt.Sprintf("/files/%s", cuid))
}

// RenameFile renames a file.
func (c *FileClient) RenameFile(ctx context.Context, cuid string, newName string) (*models.File, error) {
	req := models.RenameFileRequest{Name: newName}
	var resp struct {
		Data models.File `json:"data"`
	}
	err := c.client.Put(ctx, fmt.Sprintf("/files/%s", cuid), req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Data, nil
}
