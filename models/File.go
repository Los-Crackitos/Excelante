package models

// File : Represent json model File
type File struct {
	Sheets []*Sheet `json:"sheets"`
}
