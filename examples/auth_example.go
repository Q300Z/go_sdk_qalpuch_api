package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/clients"
	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/models"
)

func main() {
	baseURL := "http://localhost:8080/v1" // Replace with your API base URL

	// Create a new API client
	client := clients.NewClient(baseURL, "") // No token initially

	// 1. Register a new user
	registerReq := models.RegisterRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}
	registerResp, err := client.Auth.Register(context.Background(), registerReq)
	if err != nil {
		log.Fatalf("Failed to register user: %v", err)
	}
	fmt.Printf("User registered: %+v\n", registerResp.Data.User)
	fmt.Printf("Received token on registration: %s\n", registerResp.Data.Token)

	// Use the token from registration for subsequent requests
	client.Token = registerResp.Data.Token

	// 2. Login with the new user (optional, as we already have a token)
	loginReq := models.LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}
	loginResp, err := client.Auth.Login(context.Background(), loginReq)
	if err != nil {
		log.Fatalf("Failed to login: %v", err)
	}
	fmt.Printf("Login successful. Token: %s\n", loginResp.Data.Token)

	// Set the obtained token for subsequent requests
	client.Token = loginResp.Data.Token

	// 3. Get all users (requires admin role, assuming testuser is admin or API allows it)
	users, err := client.Users.GetUsers(context.Background())
	if err != nil {
		log.Fatalf("Failed to get users: %v", err)
	}
	fmt.Printf("Users: %+v\n", users)

	// 4. Change password
	changePasswordReq := models.ChangePasswordRequest{
		OldPassword: "password123",
		NewPassword: "newpassword123",
	}
	err = client.Auth.ChangePassword(context.Background(), changePasswordReq)
	if err != nil {
		log.Fatalf("Failed to change password: %v", err)
	}
	fmt.Println("Password changed successfully.")

	// 5. Logout
	logoutReq := models.LogoutRequest{
		RefreshToken: loginResp.Data.RefreshToken, // Assuming refresh token is returned on login
	}
	err = client.Auth.Logout(context.Background(), logoutReq)
	if err != nil {
		log.Fatalf("Failed to logout: %v", err)
	}
	fmt.Println("Logged out successfully.")
}
