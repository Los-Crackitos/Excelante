package models

// Header : Represent a json model Header
type Header struct {
	Orientation string    `json:"orientation,omitempty"` // Can be "col" or "row", default "col"
	Name        string    `json:"name"`
	Style       *Style    `json:"style,omitempty"`
	SubHeader   []*Header `json:"sub_headers,omitempty"`
	Cells       []*Cell   `json:"cells,omitempty"`
}
