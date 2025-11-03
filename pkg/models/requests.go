package models

type CreateTaskRequest struct {
	FileID           string      `json:"fileId"`
	Config           interface{} `json:"config,omitempty"`
	PredefinedTaskID *string     `json:"predefinedTaskId,omitempty"`
}

type RenameFileRequest struct {
	Name string `json:"name"`
}

type CreateWorkerRequest struct {
	Name         string   `json:"name"`
	Capabilities []string `json:"capabilities"`
}

type RegisterWorkerRequest struct {
	Token string `json:"token"`
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type UpdateUserRequest struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
	Role  *string `json:"role,omitempty"`
}

type UpdateTaskStatusRequest struct {
	Status        TaskStatus `json:"status"`
	StatusMessage string     `json:"message"`
}

type UpdateTaskRequest struct {
	Status       *TaskStatus `json:"status,omitempty"`
	ResultFileID *string     `json:"resultFileId,omitempty"`
}
