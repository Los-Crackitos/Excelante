package models

// Table represent json model column or row
type Table struct {
	Name  string  `json:"name"`
	Style *Style  `json:"style,omitempty"`
	Cells []*Cell `json:"cells,omitempty"`
}

// Column represent json model Column
type Column struct {
	Table
	Columns []*Column `json:"column,omitempty"`
}

// Row represent json model Row
type Row struct {
	Table
	Rows []*Row `json:"row,omitempty"`
}
