package models

//Alignment : Represent a cell style object
type Alignment struct {
	Horizontal string `json:"horizontal,omitempty"`
	Vertical   string `json:"vertical,omitempty"`
	WrapText   bool   `json:"wrap_text,omitempty"`
}
