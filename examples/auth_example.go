package main

import (
	"context"
	"fmt"
	"go_sdk_qalpuch_api/pkg/clients"
	"go_sdk_qalpuch_api/pkg/models"
	"log"
)

func main() {
	// Initialize the API client
	baseURL := "http://localhost:8080"           // Replace with your API base URL
	client := clients.NewUserClient(baseURL, "") // No token initially for login/register

	// --- User Registration Example ---
	fmt.Println("\n--- User Registration ---")
	registerReq := models.RegisterRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	registerResp, err := client.Auth.Register(context.Background(), registerReq)
	if err != nil {
		log.Printf("User registration failed: %v", err)
	} else {
		fmt.Printf("User registered successfully: %s (ID: %d)\n", registerResp.Data.Name, registerResp.Data.ID)
	}

	// --- User Login Example ---
	fmt.Println("\n--- User Login ---")
	loginReq := models.LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	loginResp, err := client.Auth.Login(context.Background(), loginReq)
	if err != nil {
		log.Fatalf("User login failed: %v", err)
	}

	fmt.Printf("User logged in successfully. Token: %s\n", loginResp.Data.Token)

	// Update the client with the new token for authenticated requests
	client.Token = loginResp.Data.Token

	// --- Get Users Example (requires admin role, assuming 'testuser' is not admin) ---
	fmt.Println("\n--- Get Users (expecting error if not admin) ---")
	users, err := client.Users.GetUsers(context.Background())
	if err != nil {
		log.Printf("Failed to get users: %v", err)
	} else {
		fmt.Println("Users:")
		for _, user := range users {
			fmt.Printf("- ID: %d, Name: %s, Email: %s, Role: %s\n", user.ID, user.Name, user.Email, user.Role)
		}
	}

	// --- Logout Example ---
	fmt.Println("\n--- User Logout ---")
	err = client.Auth.Logout(context.Background(), models.LogoutRequest{RefreshToken: "your_refresh_token"})
	if err != nil {
		log.Printf("User logout failed: %v", err)
	} else {
		fmt.Println("User logged out successfully.")
	}
}
