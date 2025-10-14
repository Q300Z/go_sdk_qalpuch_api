package models

import "time"

type Worker struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Token     *string   `json:"token,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
