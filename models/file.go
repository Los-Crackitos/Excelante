package models

// File represent  json model File
type File struct {
	Filename string   `json:"filename,omitempty"`
	Sheets   []*Sheet `json:"sheets"`
}
