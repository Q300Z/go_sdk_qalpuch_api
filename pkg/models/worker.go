package models

import "time"

type Worker struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Token        *string      `json:"token,omitempty"`
	Status       WorkerStatus `json:"status,omitempty"`
	Capabilities []WorkerType `json:"capabilities,omitempty"`
	UserID       int          `json:"userId,omitempty"`
	Tasks        *[]Task      `json:"tasks,omitempty"`
	CreatedAt    *time.Time   `json:"createdAt,omitempty"`
	UpdatedAt    *time.Time   `json:"updatedAt,omitempty"`
}
