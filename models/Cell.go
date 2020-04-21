package models

// Cell : Represent  json model Data
type Cell struct {
	Value interface{} `json:"value"`
	Style *Style      `json:"style,omitempty"`
}

// Style represent json model Data
type Style struct {
	Border     []*Border   `json:"border,omitempty"`
	Fill       *Fill       `json:"fill,omitempty"`
	Font       *Font       `json:"font,omitempty"`
	Alignment  *Alignment  `json:"alignment,omitempty"`
	Protection *Protection `json:"protection,omitempty"`
}

//Protection represent a json cell style model
type Protection struct {
	Hidden bool `json:"hidden,omitempty"`
	Locked bool `json:"locked,omitempty"`
}

//Font represent a json cell style model
type Font struct {
	Bold      bool    `json:"bold,omitempty"`
	Italic    bool    `json:"italic,omitempty"`
	Underline string  `json:"underline,omitempty"`
	Family    string  `json:"family,omitempty"`
	Size      float64 `json:"size,omitempty"`
	Strike    bool    `json:"strike,omitempty"`
	Color     string  `json:"color,omitempty"`
}

//Fill represent a json cell style model
type Fill struct {
	Type    string   `json:"type,omitempty"`
	Pattern int      `json:"pattern,omitempty"`
	Color   []string `json:"color,omitempty"`
	Shading int      `json:"shading,omitempty"`
}

//Border represent a json cell style model
type Border struct {
	Type  []string `json:"type,omitempty"`
	Color string   `json:"color,omitempty"`
	Style int      `json:"style,omitempty"`
}

//Alignment represent a json cell style model
type Alignment struct {
	Horizontal string `json:"horizontal,omitempty"`
	Vertical   string `json:"vertical,omitempty"`
	WrapText   bool   `json:"wrap_text,omitempty"`
}
