package models

// Header represent json model column or row
type Header struct {
	Name  string  `json:"name"`
	Style *Style  `json:"style,omitempty"`
	Cells []*Cell `json:"cells,omitempty"`
}

// Column represent json model Column
type Column struct {
	Header
	Columns []*Column `json:"column,omitempty"`
}

// Row represent json model Row
type Row struct {
	Header
	Rows []*Row `json:"row,omitempty"`
}
