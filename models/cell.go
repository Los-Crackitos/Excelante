package models

// Cell struct contain all the cell data used to generate a cell
type Cell struct {
	Value interface{} `json:"value"`
	Style *Style      `json:"style,omitempty"`
}

// Style represent JSON model Data
type Style struct {
	Border     []*Border   `json:"border,omitempty"`
	Fill       *Fill       `json:"fill,omitempty"`
	Font       *Font       `json:"font,omitempty"`
	Alignment  *Alignment  `json:"alignment,omitempty"`
	Protection *Protection `json:"protection,omitempty"`
}

//Protection represent a JSON cell style model
type Protection struct {
	Hidden bool `json:"hidden,omitempty"`
	Locked bool `json:"locked,omitempty"`
}

//Font represent a JSON cell style model
type Font struct {
	Bold      bool    `json:"bold,omitempty"`
	Italic    bool    `json:"italic,omitempty"`
	Underline string  `json:"underline,omitempty"`
	Family    string  `json:"family,omitempty"`
	Size      float64 `json:"size,omitempty"`
	Strike    bool    `json:"strike,omitempty"`
	Color     string  `json:"color,omitempty"`
}

// Fill represent a JSON cell style model
type Fill struct {
	Type    string   `json:"type,omitempty"`
	Pattern int      `json:"pattern,omitempty"`
	Color   []string `json:"color,omitempty"`
	Shading int      `json:"shading,omitempty"`
}

// Border struct contain all the border data used to stylize a border of a cell, column or row
type Border struct {
	Type  string `json:"type,omitempty"`
	Color string `json:"color,omitempty"`
	Style int    `json:"style,omitempty"`
}

// Alignment struct contain the text alignment details that will be used inside cell
type Alignment struct {
	Horizontal  string `json:"horizontal,omitempty"`
	Vertical    string `json:"vertical,omitempty"`
	ShrinkToFit bool   `json:"shrink_to_fit"`
	WrapText    bool   `json:"wrap_text,omitempty"`
}
