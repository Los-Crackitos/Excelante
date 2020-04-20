package models

// Cell : Represent json model Data
type Cell struct {
	Value string `json:"value"`
	Style *Style `json:"style,omitempty"`
}
