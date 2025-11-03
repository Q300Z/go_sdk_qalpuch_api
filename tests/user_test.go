package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/clients"
	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/models"
)

func TestUserClient_GetUsers(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/users" {
			t.Errorf("Expected to request '/v1/users', got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		response := struct {
			Success bool          `json:"success"`
			Data    []models.User `json:"data"`
		}{
			Success: true,
			Data:    []models.User{{ID: 1, Name: "User 1"}, {ID: 2, Name: "User 2"}},
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	c := clients.NewClient(server.URL+"/v1", "test_token")

	users, err := c.Users.GetUsers(context.Background())
	if err != nil {
		t.Fatalf("GetUsers failed: %v", err)
	}

	if len(users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(users))
	}
}

func TestUserClient_GetUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/users/1" {
			t.Errorf("Expected to request '/v1/users/1', got %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		response := struct {
			Success bool        `json:"success"`
			Data    models.User `json:"data"`
		}{
			Success: true,
			Data:    models.User{ID: 1, Name: "User 1"},
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	c := clients.NewClient(server.URL+"/v1", "test_token")

	user, err := c.Users.GetUser(context.Background(), 1)
	if err != nil {
		t.Fatalf("GetUser failed: %v", err)
	}

	if user.ID != 1 {
		t.Errorf("Expected user ID 1, got %d", user.ID)
	}
}

func TestUserClient_UpdateUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/users/1" {
			t.Errorf("Expected to request '/v1/users/1', got %s", r.URL.Path)
		}
		if r.Method != http.MethodPut {
			t.Errorf("Expected PUT request, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		response := struct {
			Success bool        `json:"success"`
			Data    models.User `json:"data"`
		}{
			Success: true,
			Data:    models.User{ID: 1, Name: "Updated User"},
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	c := clients.NewClient(server.URL+"/v1", "test_token")

	name := "Updated User"
	updateReq := models.UpdateUserRequest{Name: &name}
	user, err := c.Users.UpdateUser(context.Background(), 1, updateReq)
	if err != nil {
		t.Fatalf("UpdateUser failed: %v", err)
	}

	if user.Name != "Updated User" {
		t.Errorf("Expected user name 'Updated User', got '%s'", user.Name)
	}
}

func TestUserClient_DeleteCurrentUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/users/me" {
			t.Errorf("Expected to request '/v1/users/me', got %s", r.URL.Path)
		}
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE request, got %s", r.Method)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c := clients.NewClient(server.URL+"/v1", "test_token")

	err := c.Users.DeleteCurrentUser(context.Background())
	if err != nil {
		t.Fatalf("DeleteCurrentUser failed: %v", err)
	}
}

func TestUserClient_CreateUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/users" {
			t.Errorf("Expected to request '/v1/users', got %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(`{"success": true, "data": {"id": 1, "name": "new user"}}`)); err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	c := clients.NewClient(server.URL+"/v1", "test_token")

	req := models.CreateUserRequest{
		Name:  "new user",
		Email: "new@user.com",
		Role:  "user",
	}

	_, err := c.Users.CreateUser(context.Background(), req)
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}
}

func TestUserClient_SearchUsers(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/users/search" {
			t.Errorf("Expected to request '/v1/users/search', got %s", r.URL.Path)
		}
		if r.URL.Query().Get("q") != "test" {
			t.Errorf("Expected query param q=test, got %s", r.URL.Query().Get("q"))
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(`{"success": true, "data": []}`)); err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	c := clients.NewClient(server.URL+"/v1", "test_token")

	_, err := c.Users.SearchUsers(context.Background(), "test")
	if err != nil {
		t.Fatalf("SearchUsers failed: %v", err)
	}
}
