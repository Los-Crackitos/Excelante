package models

// Table struct is containing all the cells values that will be insert into a column or row
type Table struct {
	Orientation string  `json:"orientation,omitempty"`
	Cells       []*Cell `json:"cells,omitempty"`
}
