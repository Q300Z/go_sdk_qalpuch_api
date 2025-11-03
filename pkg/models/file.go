package models

import "time"

type File struct {
	ID        string    `json:"id"`
	Filename  string    `json:"filename"`
	Size      int64     `json:"size"`
	MimeType  string    `json:"mimetype"`
	Hash      string    `json:"hash"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
