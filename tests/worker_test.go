package tests

import (
	"context"
	"encoding/json"
	"go_sdk_qalpuch_api/pkg/clients"
	"go_sdk_qalpuch_api/pkg/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWorkerClient_RegisterWorker(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/worker/register" {
			t.Errorf("Expected to request '/v1/worker/register', got %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
		response := models.AuthWorkerResponse{
			Success: true,
			Message: "Worker authenticated successfully",
			Data: struct {
				Token        string `json:"token"`
				RefreshToken string `json:"refreshToken"`
				ExpiresIn    int    `json:"expiresIn"`
			}{
				Token:        "worker_jwt_token",
				RefreshToken: "worker_refresh_token",
				ExpiresIn:    3600,
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	c := clients.NewWorkerClient(server.URL, "") // No token initially for registration

	resp, err := c.Workers.RegisterWorker(context.Background(), "a1b2c3d4-e5f6-7890-1234-567890abcdef")
	if err != nil {
		t.Fatalf("RegisterWorker failed: %v", err)
	}

	if !resp.Success {
		t.Errorf("Expected success to be true, got false")
	}

	if resp.Data.Token != "worker_jwt_token" {
		t.Errorf("Expected token 'worker_jwt_token', got %s", resp.Data.Token)
	}
}

func TestWorkerClient_CreateWorker(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/worker" {
			t.Errorf("Expected to request '/v1/worker', got %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}

		var reqBody map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			t.Fatalf("Failed to decode request body: %v", err)
		}
		if reqBody["name"] != "test-worker" {
			t.Errorf("Expected name 'test-worker', got '%s'", reqBody["name"])
		}

		w.WriteHeader(http.StatusCreated)
		response := struct {
			Success bool          `json:"success"`
			Data    models.Worker `json:"data"`
		}{
			Success: true,
			Data:    models.Worker{ID: "new-worker-id", Name: "test-worker"},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	c := clients.NewUserClient(server.URL, "admin_token")

	worker, err := c.Workers.CreateWorker(context.Background(), "test-worker", []string{"video", "image"})
	if err != nil {
		t.Fatalf("CreateWorker failed: %v", err)
	}

	if worker.ID != "new-worker-id" {
		t.Errorf("Expected worker ID 'new-worker-id', got '%s'", worker.ID)
	}
}
