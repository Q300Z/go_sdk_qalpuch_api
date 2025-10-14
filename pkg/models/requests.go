package models

// CreateTaskRequest model for creating a new task
type CreateTaskRequest struct {
	FileID string      `json:"file_id"`
	Config interface{} `json:"config"`
}

// UpdateTaskStatusRequest model for updating a task's status
type UpdateTaskStatusRequest struct {
	Status        TaskStatus `json:"status"`
	StatusMessage string     `json:"status_message"`
}

// CreateUserRequest model for an admin to create a new user
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
