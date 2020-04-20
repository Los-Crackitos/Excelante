package models

//Font : Represent a json cell style model
type Font struct {
	Bold      bool    `json:"bold,omitempty"`
	Italic    bool    `json:"italic,omitempty"`
	Underline string  `json:"underline,omitempty"`
	Family    string  `json:"family,omitempty"`
	Size      float64 `json:"size,omitempty"`
	Strike    bool    `json:"strike,omitempty"`
	Color     string  `json:"color,omitempty"`
}
