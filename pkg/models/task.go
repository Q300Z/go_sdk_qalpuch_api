package models

import "time"

type TaskStatus string

const (
	TaskStatusProcessing TaskStatus = "processing"
)

type Task struct {
	ID           string     `json:"id"`
	Config       string     `json:"config"`
	Type         WorkerType `json:"type"`
	Status       TaskStatus `json:"status"`
	SourceFileID string     `json:"sourceFileId"`
	ResultFileID *string    `json:"result_file_id,omitempty"`
	WorkerID     *string    `json:"workerId,omitempty"`
	Logs         []Log      `json:"logs,omitempty"`
	Message      *string    `json:"message,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

type Log struct {
	ID         string     `json:"id"`
	TaskID     string     `json:"taskId"`
	TaskStatus TaskStatus `json:"task_status"`
	Message    string     `json:"message"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
