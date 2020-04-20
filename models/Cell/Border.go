package models

//Border : Represent a json cell style model
type Border struct {
	Type  []string `json:"type,omitempty"`
	Color string   `json:"color,omitempty"`
	Style int      `json:"style,omitempty"`
}
