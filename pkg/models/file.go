package models

type File struct {
	ID       string `json:"id"`
	Filename string `json:"filename"`
	Path     string `json:"path"`
	Size     int    `json:"size"`
	MimeType string `json:"mimeType"`
}
