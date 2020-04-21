package models

// File represent  json model File
type File struct {
	Sheets []*Sheet `json:"sheets"`
}
