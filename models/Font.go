package models

type Font struct {
	Bold      bool    `json:"bold"`
	Italic    bool    `json:"italic"`
	Underline string  `json:"underline"`
	Family    string  `json:"family"`
	Size      float64 `json:"size"`
	Strike    bool    `json:"strike"`
	Color     string  `json:"color"`
}
