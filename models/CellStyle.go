package models

//CellStyle : Represent a cell style
type CellStyle struct {
	Border     []Border    `json:"border,omitempty"`
	Fill       Fill        `json:"fill,omitempty"`
	Font       *Font       `json:"font,omitempty"`
	Alignment  *Alignment  `json:"alignment,omitempty"`
	Protection *Protection `json:"protection,omitempty"`
}
