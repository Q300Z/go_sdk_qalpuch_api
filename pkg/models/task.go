package models

import "time"

type TaskStatus string

const (
	TaskStatusPending    TaskStatus = "pending"
	TaskStatusProcessing TaskStatus = "processing"
	TaskStatusCompleted  TaskStatus = "completed"
	TaskStatusFailed     TaskStatus = "failed"
)

type Task struct {
	ID           string     `json:"id"`
	Config       string     `json:"config"`
	Status       TaskStatus `json:"status"`
	SourceFileID string     `json:"source_file_id"`
	ResultFileID *string    `json:"result_file_id"`
	Logs         []Log      `json:"logs"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

type Log struct {
	ID         string    `json:"id"`
	TaskStatus string    `json:"task_status"`
	Message    string    `json:"message"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
