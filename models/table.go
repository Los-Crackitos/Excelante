package models

// ColOrRowValues struct is containing all the cells values that will be insert into a column or row
type ColOrRowValues struct {
	Orientation string  `json:"orientation,omitempty"`
	Cells       []*Cell `json:"cells,omitempty"`
}
