package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/errors"
	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/models"
	"io"
	"mime/multipart"
	"net/http"
)

// FileClient handles file-related API requests.
type FileClient struct {
	client *Client
}

// UploadFile uploads a file.
func (c *FileClient) UploadFile(ctx context.Context, name string, file []byte) (*models.File, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", name)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, bytes.NewReader(file))
	if err != nil {
		return nil, err
	}

	writer.Close()

	url := fmt.Sprintf("%s/v1/files/upload", c.client.BaseURL)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", writer.FormDataContentType())
	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.client.Token))

	resp, err := c.client.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.NewAPIErrorFromResponse(resp)
	}

	var fileResp struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    models.File `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&fileResp); err != nil {
		return nil, err
	}

	return &fileResp.Data, nil
}

// GetFileMetadata retrieves a file's metadata.
func (c *FileClient) GetFileMetadata(ctx context.Context, cuid string) (*models.File, error) {
	url := fmt.Sprintf("%s/v1/files/%s", c.client.BaseURL, cuid)
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

	var fileResp struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    models.File `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&fileResp); err != nil {
		return nil, err
	}

	return &fileResp.Data, nil
}

// DownloadFile downloads a file.
func (c *FileClient) DownloadFile(ctx context.Context, cuid string) ([]byte, error) {
	url := fmt.Sprintf("%s/v1/files/%s/download", c.client.BaseURL, cuid)
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

	return io.ReadAll(resp.Body)
}

// ListUserFiles lists a user's files.
func (c *FileClient) ListUserFiles(ctx context.Context) ([]models.File, error) {
	url := fmt.Sprintf("%s/v1/files", c.client.BaseURL)
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

	var filesResp struct {
		Success bool          `json:"success"`
		Message string        `json:"message"`
		Data    []models.File `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&filesResp); err != nil {
		return nil, err
	}

	return filesResp.Data, nil
}

// DeleteFile deletes a file.
func (c *FileClient) DeleteFile(ctx context.Context, cuid string) error {
	url := fmt.Sprintf("%s/v1/files/%s", c.client.BaseURL, cuid)
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

// RenameFile renames a file.
func (c *FileClient) RenameFile(ctx context.Context, cuid string, newName string) (*models.File, error) {
	requestBody := map[string]string{"name": newName}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/v1/files/%s", c.client.BaseURL, cuid)
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

	var fileResp struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    models.File `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&fileResp); err != nil {
		return nil, err
	}

	return &fileResp.Data, nil
}
