package models

import "time"

type Worker struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Token        *string      `json:"token,omitempty"`
	Status       WorkerStatus `json:"status"`
	Capabilities []WorkerType `json:"capabilities"`
	UserID       int          `json:"userId"`
	Tasks        []Task       `json:"tasks,omitempty"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}
