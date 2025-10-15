package tests

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/clients"
	sdkerrors "github.com/Q300Z/go_sdk_qalpuch_api/pkg/errors"
	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTaskClient_GetUserTasks(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/tasks" {
			t.Errorf("Expected to request '/v1/tasks', got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		response := struct {
			Success bool          `json:"success"`
			Data    []models.Task `json:"data"`
		}{
			Success: true,
			Data:    []models.Task{{ID: "task1"}, {ID: "task2"}},
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	c := clients.NewUserClient(server.URL, "test_token")

	tasks, err := c.Tasks.GetUserTasks(context.Background())
	if err != nil {
		t.Fatalf("GetUserTasks failed: %v", err)
	}

	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}
}

func TestTaskClient_CreateTask(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/tasks" {
			t.Errorf("Expected to request '/v1/tasks', got %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		w.WriteHeader(http.StatusCreated)
		response := struct {
			Success bool        `json:"success"`
			Data    models.Task `json:"data"`
		}{
			Success: true,
			Data:    models.Task{ID: "new-task-id"},
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	c := clients.NewUserClient(server.URL, "test_token")

	req := models.CreateTaskRequest{FileID: "file-id", Config: "{}"}
	task, err := c.Tasks.CreateTask(context.Background(), req)
	if err != nil {
		t.Fatalf("CreateTask failed: %v", err)
	}

	if task.ID != "new-task-id" {
		t.Errorf("Expected task ID 'new-task-id', got '%s'", task.ID)
	}
}

func TestTaskClient_DeleteTask(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/tasks/task-to-delete" {
			t.Errorf("Expected to request '/v1/tasks/task-to-delete', got %s", r.URL.Path)
		}
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE request, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c := clients.NewUserClient(server.URL, "test_token")

	err := c.Tasks.DeleteTask(context.Background(), "task-to-delete")
	if err != nil {
		t.Fatalf("DeleteTask failed: %v", err)
	}
}

func TestTaskClient_GetPendingTask(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/tasks/pending" {
			t.Errorf("Expected to request '/v1/tasks/pending', got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		response := struct {
			Success bool        `json:"success"`
			Data    models.Task `json:"data"`
		}{
			Success: true,
			Data:    models.Task{ID: "pending-task-id"},
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	c := clients.NewWorkerClient(server.URL, "test_token")

	task, err := c.Tasks.GetPendingTask(context.Background())
	if err != nil {
		t.Fatalf("GetPendingTask failed: %v", err)
	}

	if task.ID != "pending-task-id" {
		t.Errorf("Expected task ID 'pending-task-id', got '%s'", task.ID)
	}
}

func TestTaskClient_UpdateTaskStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/tasks/task-to-update" {
			t.Errorf("Expected to request '/v1/tasks/task-to-update', got %s", r.URL.Path)
		}
		if r.Method != http.MethodPatch {
			t.Errorf("Expected PATCH request, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c := clients.NewWorkerClient(server.URL, "test_token")

	req := models.UpdateTaskStatusRequest{Status: models.TaskStatusProcessing}
	err := c.Tasks.UpdateTaskStatus(context.Background(), "task-to-update", req)
	if err != nil {
		t.Fatalf("UpdateTaskStatus failed: %v", err)
	}
}

func TestTaskClient_UploadTaskResult(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/tasks/clv9p8qjk000108l9e7g3f2a1/result" {
			t.Errorf("Expected to request '/v1/tasks/clv9p8qjk000108l9e7g3f2a1/result', got %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c := clients.NewWorkerClient(server.URL, "test_token")

	resultContent := []byte("task result content")

	err := c.Tasks.UploadTaskResult(context.Background(), "clv9p8qjk000108l9e7g3f2a1", resultContent)
	if err != nil {
		t.Fatalf("UploadTaskResult failed: %v", err)
	}
}

func TestTaskClient_BuildAndExecute(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/tasks" {
			t.Errorf("Expected to request '/v1/tasks', got %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}

		var req models.CreateTaskRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("Failed to decode request body: %v", err)
		}

		if req.FileID != "test-file-id" {
			t.Errorf("Expected file ID 'test-file-id', got '%s'", req.FileID)
		}

		// Check the config
		var videoConf models.VideoConversionConfig
		configBytes, err := json.Marshal(req.Config)
		if err != nil {
			t.Fatalf("Failed to marshal config map: %v", err)
		}
		if err := json.Unmarshal(configBytes, &videoConf); err != nil {
			t.Fatalf("Failed to unmarshal config from request: %v", err)
		}
		if videoConf.Codec != "h264" {
			t.Errorf("Expected codec 'h264', got '%s'", videoConf.Codec)
		}

		w.WriteHeader(http.StatusCreated)
		response := struct {
			Success bool        `json:"success"`
			Data    models.Task `json:"data"`
		}{
			Success: true,
			Data:    models.Task{ID: "new-builder-task"},
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	c := clients.NewUserClient(server.URL, "test_token")

	videoConf := models.NewVideoConfig().WithCodec("h264")

	task, err := c.Tasks.Build("test-file-id").WithVideoConfig(*videoConf).Execute(context.Background())

	if err != nil {
		t.Fatalf("BuildAndExecute failed: %v", err)
	}

	if task.ID != "new-builder-task" {
		t.Errorf("Expected task ID 'new-builder-task', got '%s'", task.ID)
	}
}

func TestTaskClient_GetUserTasks_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	c := clients.NewUserClient(server.URL, "test_token")

	_, err := c.Tasks.GetUserTasks(context.Background())
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if !errors.Is(err, sdkerrors.ErrNotFound) {
		t.Errorf("Expected error to be of type ErrNotFound, got %T", err)
	}
}
