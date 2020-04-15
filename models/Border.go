package models

//Border : Represent a cell style object
type Border struct {
	Type  []string `json:"type,omitempty"`
	Color string   `json:"color,omitempty"`
	Style int      `json:"style,omitempty"`
}
