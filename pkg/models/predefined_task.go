package models

import "time"

type CreatePredefinedTaskRequest struct {
	Name        string                 `json:"name"`
	Description *string                `json:"description,omitempty"`
	Config      map[string]interface{} `json:"config"`
}

type UpdatePredefinedTaskRequest struct {
	Name        *string                `json:"name,omitempty"`
	Description *string                `json:"description,omitempty"`
	Config      map[string]interface{} `json:"config,omitempty"`
}

type PredefinedTask struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	Config      interface{} `json:"config"` // Stored as JSON string
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
	AdminID     int       `json:"adminId"`
}
