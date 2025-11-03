package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/clients"
	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/models"
)

func StringPtr(s string) *string {
	return &s
}

func setupPredefinedTaskTestServer(handler http.HandlerFunc) (*httptest.Server, *clients.Client) {
	server := httptest.NewServer(handler)
	client := clients.NewClient(server.URL+"/v1", "test_token")
	return server, client
}

func TestPredefinedTaskClient_CreatePredefinedTask_Success(t *testing.T) {
	server, client := setupPredefinedTaskTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/predefined-tasks" {
			t.Errorf("Expected to request '/v1/predefined-tasks', got %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		w.WriteHeader(http.StatusCreated)
		now := time.Now()
		response := models.PredefinedTask{
			ID:          "test-cuid",
			Name:        "Test Predefined Task",
			Description: StringPtr("A test description"),
			Config:      `{"type":"video"}`,
			CreatedAt:   &now,
			UpdatedAt:   &now,
			AdminID:     1,
		}
		var data interface{} = response
		if err := json.NewEncoder(w).Encode(models.APIResponse{Success: true, Data: &data}); err != nil {
			t.Fatal(err)
		}
	})
	defer server.Close()

	req := models.CreatePredefinedTaskRequest{
		Name:        "Test Predefined Task",
		Description: StringPtr("A test description"),
		Config:      map[string]interface{}{"type": "video"},
	}

	resp, err := client.PredefinedTasks.CreatePredefinedTask(context.Background(), req)
	if err != nil {
		t.Fatalf("CreatePredefinedTask failed: %v", err)
	}

	if resp.ID != "test-cuid" {
		t.Errorf("Expected ID 'test-cuid', got '%s'", resp.ID)
	}
}

func TestPredefinedTaskClient_GetPredefinedTasks_Success(t *testing.T) {
	server, client := setupPredefinedTaskTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/predefined-tasks" {
			t.Errorf("Expected to request '/v1/predefined-tasks', got %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		response := []models.PredefinedTask{
			{ID: "id1", Name: "Task1", Config: `{"type":"image"}`},
			{ID: "id2", Name: "Task2", Config: `{"type":"audio"}`},
		}
		var data interface{} = response
		if err := json.NewEncoder(w).Encode(models.APIResponse{Success: true, Data: &data}); err != nil {
			t.Fatal(err)
		}
	})
	defer server.Close()

	resp, err := client.PredefinedTasks.GetPredefinedTasks(context.Background())
	if err != nil {
		t.Fatalf("GetPredefinedTasks failed: %v", err)
	}

	if len(resp) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(resp))
	}
}

func TestPredefinedTaskClient_GetPredefinedTaskByID_Success(t *testing.T) {
	taskID := "test-id"
	server, client := setupPredefinedTaskTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != fmt.Sprintf("/v1/predefined-tasks/%s", taskID) {
			t.Errorf("Expected to request '/v1/predefined-tasks/%s', got %s", taskID, r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		response := models.PredefinedTask{
			ID:     taskID,
			Name:   "Test Task",
			Config: `{"type":"video"}`,
		}
		var data interface{} = response
		if err := json.NewEncoder(w).Encode(models.APIResponse{Success: true, Data: &data}); err != nil {
			t.Fatal(err)
		}
	})
	defer server.Close()

	resp, err := client.PredefinedTasks.GetPredefinedTaskByID(context.Background(), taskID)
	if err != nil {
		t.Fatalf("GetPredefinedTaskByID failed: %v", err)
	}

	if resp.ID != taskID {
		t.Errorf("Expected ID '%s', got '%s'", taskID, resp.ID)
	}
}

func TestPredefinedTaskClient_UpdatePredefinedTask_Success(t *testing.T) {
	taskID := "update-id"
	server, client := setupPredefinedTaskTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != fmt.Sprintf("/v1/predefined-tasks/%s", taskID) {
			t.Errorf("Expected to request '/v1/predefined-tasks/%s', got %s", taskID, r.URL.Path)
		}
		if r.Method != http.MethodPut {
			t.Errorf("Expected PUT request, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		newName := "Updated Name"
		response := models.PredefinedTask{
			ID:     taskID,
			Name:   newName,
			Config: `{"type":"video"}`,
		}
		var data interface{} = response
		if err := json.NewEncoder(w).Encode(models.APIResponse{Success: true, Data: &data}); err != nil {
			t.Fatal(err)
		}
	})
	defer server.Close()

	newName := "Updated Name"
	req := models.UpdatePredefinedTaskRequest{Name: &newName}

	resp, err := client.PredefinedTasks.UpdatePredefinedTask(context.Background(), taskID, req)
	if err != nil {
		t.Fatalf("UpdatePredefinedTask failed: %v", err)
	}

	if resp.Name != newName {
		t.Errorf("Expected name '%s', got '%s'", newName, resp.Name)
	}
}

func TestPredefinedTaskClient_DeletePredefinedTask_Success(t *testing.T) {
	taskID := "delete-id"
	server, client := setupPredefinedTaskTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != fmt.Sprintf("/v1/predefined-tasks/%s", taskID) {
			t.Errorf("Expected to request '/v1/predefined-tasks/%s', got %s", taskID, r.URL.Path)
		}
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE request, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(models.APIResponse{Success: true, Message: "Deleted"}); err != nil {
			t.Fatal(err)
		}
	})
	defer server.Close()

	err := client.PredefinedTasks.DeletePredefinedTask(context.Background(), taskID)
	if err != nil {
		t.Fatalf("DeletePredefinedTask failed: %v", err)
	}
}
