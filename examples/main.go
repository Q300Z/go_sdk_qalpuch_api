package main

import (
	"context"
	"fmt"
	"go_sdk_qalpuch_api/pkg/clients"
	"go_sdk_qalpuch_api/pkg/models"
	"log"
)

func main() {
	// 1. Initialize the API client
	// Replace with your API base URL. For local development, it's often http://localhost:8080.
	baseURL := "http://localhost:8080"
	client := clients.NewUserClient(baseURL, "") // No token is needed for registration or login.

	fmt.Println("--- 1. User Registration ---")
	registerReq := models.RegisterRequest{
		Username: "example-user",
		Email:    "user@example.com",
		Password: "a-strong-password",
	}
	registerResp, err := client.Auth.Register(context.Background(), registerReq)
	if err != nil {
		// This might fail if the user already exists, which is expected on subsequent runs.
		log.Printf("User registration failed (this is expected if user already exists): %v", err)
	} else {
		fmt.Printf("User registered successfully: %s (ID: %d)\n", registerResp.Data.Name, registerResp.Data.ID)
	}

	fmt.Println("\n--- 2. User Login ---")
	loginReq := models.LoginRequest{
		Email:    "user@example.com",
		Password: "a-strong-password",
	}
	loginResp, err := client.Auth.Login(context.Background(), loginReq)
	if err != nil {
		log.Fatalf("User login failed: %v", err)
	}
	fmt.Printf("User logged in successfully.\n")

	// Set the token on the client for all subsequent authenticated requests.
	client.Token = loginResp.Data.Token

	fmt.Println("\n--- 3. File Upload ---")
	// Create a dummy file to upload.
	dummyContent := []byte("This is the content of my file.")
	fileName := "my-test-file.txt"
	
	file, err := client.Files.UploadFile(context.Background(), fileName, dummyContent)
	if err != nil {
		log.Fatalf("File upload failed: %v", err)
	}
	fmt.Printf("File uploaded successfully. File ID: %s\n", file.ID)

	fmt.Println("\n--- 4. Task Creation ---")
	// Create a task to convert the uploaded file.
	// This example uses a simple video conversion configuration.
	videoConfig := models.NewVideoConfig().WithCodec("h264").WithBitrate(1000)

	task, err := client.Tasks.Build(file.ID).WithVideoConfig(*videoConfig).Execute(context.Background())
	if err != nil {
		log.Fatalf("Task creation failed: %v", err)
	}
	fmt.Printf("Task created successfully. Task ID: %s\n", task.ID)

	fmt.Println("\n--- 5. Get User's Tasks ---")
	tasks, err := client.Tasks.GetUserTasks(context.Background())
	if err != nil {
		log.Fatalf("Failed to get user tasks: %v", err)
	}
	fmt.Printf("Successfully retrieved %d tasks.\n", len(tasks))
	for _, t := range tasks {
		fmt.Printf("- Task ID: %s, Status: %s\n", t.ID, t.Status)
	}
}
