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

func TestAuthClient_Login(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/login" {
			t.Errorf("Expected to request '/v1/login', got %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		response := models.LoginResponse{
			Success: true,
			Message: "Login successful",
			Data: struct {
				Token        string `json:"token"`
				RefreshToken string `json:"refreshToken"`
			}{
				Token:        "fake_jwt_token",
				RefreshToken: "fake_refresh_token",
			},
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	c := clients.NewUserClient(server.URL, "")

	req := models.LoginRequest{
		Email:    "test@example.com",
		Password: "password",
	}

	resp, err := c.Auth.Login(context.Background(), req)
	if err != nil {
		t.Fatalf("Login failed: %v", err)
	}

	if !resp.Success {
		t.Error("Expected login to be successful")
	}

	if resp.Data.Token != "fake_jwt_token" {
		t.Errorf("Expected token 'fake_jwt_token', got '%s'", resp.Data.Token)
	}
}

func TestAuthClient_Register(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/register" {
			t.Errorf("Expected to request '/v1/register', got %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		w.WriteHeader(http.StatusCreated)
		response := models.RegisterResponse{
			Success: true,
			Message: "User registered successfully",
			Data:    models.User{ID: 1, Name: "testuser"},
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	c := clients.NewUserClient(server.URL, "")

	req := models.RegisterRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	resp, err := c.Auth.Register(context.Background(), req)
	if err != nil {
		t.Fatalf("Register failed: %v", err)
	}

	if !resp.Success {
		t.Error("Expected registration to be successful")
	}

	if resp.Data.ID != 1 {
		t.Errorf("Expected user ID 1, got %d", resp.Data.ID)
	}
}

func TestAuthClient_Logout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/logout" {
			t.Errorf("Expected to request '/v1/logout', got %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		authHeader := r.Header.Get("Authorization")
		expectedHeader := "Bearer test_token"
		if authHeader != expectedHeader {
			t.Errorf("Expected Authorization header '%s', got '%s'", expectedHeader, authHeader)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c := clients.NewUserClient(server.URL, "test_token")

	req := models.LogoutRequest{
		RefreshToken: "fake_refresh_token",
	}

	err := c.Auth.Logout(context.Background(), req)
	if err != nil {
		t.Fatalf("Logout failed: %v", err)
	}
}

func TestAuthClient_ChangePassword(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/change-password" {
			t.Errorf("Expected to request '/v1/change-password', got %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		authHeader := r.Header.Get("Authorization")
		expectedHeader := "Bearer test_token"
		if authHeader != expectedHeader {
			t.Errorf("Expected Authorization header '%s', got '%s'", expectedHeader, authHeader)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c := clients.NewUserClient(server.URL, "test_token")

	req := models.ChangePasswordRequest{
		OldPassword: "old_password",
		NewPassword: "new_password",
	}

	err := c.Auth.ChangePassword(context.Background(), req)
	if err != nil {
		t.Fatalf("ChangePassword failed: %v", err)
	}
}

func TestAuthClient_RefreshToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/refresh" {
			t.Errorf("Expected to request '/v1/refresh', got %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		response := models.RefreshResponse{
			Success: true,
			Data: struct {
				Token        string `json:"token"`
				RefreshToken string `json:"refreshToken"`
				ExpiresIn    int    `json:"expiresIn"`
			}{
				Token:        "new_access_token",
				RefreshToken: "new_refresh_token",
				ExpiresIn:    3600,
			},
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	c := clients.NewUserClient(server.URL, "")

	req := models.RefreshTokenRequest{
		RefreshToken: "fake_refresh_token",
	}

	resp, err := c.Auth.RefreshToken(context.Background(), req)
	if err != nil {
		t.Fatalf("RefreshToken failed: %v", err)
	}

	if resp.Data.Token != "new_access_token" {
		t.Errorf("Expected token 'new_access_token', got '%s'", resp.Data.Token)
	}
}
