package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/clients"
	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/models"
)

func TestFileClient_UploadFile(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/files/upload" {
			t.Errorf("Expected to request '/v1/files/upload', got %s", r.URL.Path)
		}
		if !strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data") {
			t.Errorf("Expected Content-Type multipart/form-data, got %s", r.Header.Get("Content-Type"))
		}
		w.WriteHeader(http.StatusOK)
		response := struct {
			Success bool        `json:"success"`
			Data    models.File `json:"data"`
		}{
			Success: true,
			Data:    models.File{ID: "new-file-id", Filename: "test.txt"},
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	c := clients.NewClient(server.URL+"/v1", "test_token")

	fileContents := []byte("this is a test file")
	_, err := c.Files.UploadFile(context.Background(), "test.txt", fileContents)
	if err != nil {
		t.Fatalf("UploadFile failed: %v", err)
	}
}

func TestFileClient_GetFileMetadata(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/files/test-cuid" {
			t.Errorf("Expected to request '/v1/files/test-cuid', got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		response := struct {
			Success bool        `json:"success"`
			Data    models.File `json:"data"`
		}{
			Success: true,
			Data:    models.File{ID: "test-cuid"},
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	c := clients.NewClient(server.URL+"/v1", "test_token")

	file, err := c.Files.GetFileMetadata(context.Background(), "test-cuid")
	if err != nil {
		t.Fatalf("GetFileMetadata failed: %v", err)
	}

	if file.ID != "test-cuid" {
		t.Errorf("Expected file ID 'test-cuid', got '%s'", file.ID)
	}
}

func TestFileClient_DownloadFile(t *testing.T) {
	fileContent := "hello world"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/files/test-cuid/download" {
			t.Errorf("Expected to request '/v1/files/test-cuid/download', got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(fileContent)); err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	c := clients.NewClient(server.URL+"/v1", "test_token")

	data, err := c.Files.DownloadFile(context.Background(), "test-cuid")
	if err != nil {
		t.Fatalf("DownloadFile failed: %v", err)
	}

	if string(data) != fileContent {
		t.Errorf("Expected file content '%s', got '%s'", fileContent, string(data))
	}
}

func TestFileClient_ListUserFiles(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/files" {
			t.Errorf("Expected to request '/v1/files', got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		response := struct {
			Success bool          `json:"success"`
			Data    []models.File `json:"data"`
		}{
			Success: true,
			Data:    []models.File{{ID: "file1"}, {ID: "file2"}},
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	c := clients.NewClient(server.URL+"/v1", "test_token")

	files, err := c.Files.ListUserFiles(context.Background())
	if err != nil {
		t.Fatalf("ListUserFiles failed: %v", err)
	}

	if len(files) != 2 {
		t.Errorf("Expected 2 files, got %d", len(files))
	}
}

func TestFileClient_DeleteFile(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/files/clvb2qabc000008l21234abcd" {
			t.Errorf("Expected to request '/v1/files/clvb2qabc000008l21234abcd', got %s", r.URL.Path)
		}
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c := clients.NewClient(server.URL+"/v1", "test_token")

	err := c.Files.DeleteFile(context.Background(), "clvb2qabc000008l21234abcd")
	if err != nil {
		t.Fatalf("DeleteFile failed: %v", err)
	}
}

func TestFileClient_RenameFile(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/files/test-cuid" {
			t.Errorf("Expected to request '/v1/files/test-cuid', got %s", r.URL.Path)
		}
		if r.Method != http.MethodPut {
			t.Errorf("Expected PUT request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(`{"success": true, "data": {"id": "test-cuid", "filename": "new-name.txt"}}`)); err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	c := clients.NewClient(server.URL+"/v1", "test_token")

	_, err := c.Files.RenameFile(context.Background(), "test-cuid", "new-name.txt")
	if err != nil {
		t.Fatalf("RenameFile failed: %v", err)
	}
}
