package models

// Style : Represent json model Data
type Style struct {
	Border     []Border    `json:"border,omitempty"`
	Fill       Fill        `json:"fill,omitempty"`
	Font       *Font       `json:"font,omitempty"`
	Alignment  *Alignment  `json:"alignment,omitempty"`
	Protection *Protection `json:"protection,omitempty"`
}
