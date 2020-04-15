package models

//Fill : Represent a cell style object
type Fill struct {
	Type    string   `json:"type,omitempty"`
	Pattern int      `json:"pattern,omitempty"`
	Color   []string `json:"color,omitempty"`
	Shading int      `json:"shading,omitempty"`
}
