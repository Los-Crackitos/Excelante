package models

// Column represent json model Column
type Column struct {
	Name   string    `json:"name"`
	Style  *Style    `json:"style,omitempty"`
	Column []*Column `json:"column,omitempty"`
	Cells  []*Cell   `json:"cells,omitempty"`
}

// Row represent json model Row
type Row struct {
	Name  string  `json:"name"`
	Style *Style  `json:"style,omitempty"`
	Row   []*Row  `json:"row,omitempty"`
	Cells []*Cell `json:"cells,omitempty"`
}
