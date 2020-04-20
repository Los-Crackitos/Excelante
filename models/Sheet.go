package models

// Sheet represent json model Sheet
type Sheet struct {
	Name        string  `json:"name"`
	Orientation string  `json:"orientation,omitempty"` // Default : landscape
	Items       []*Item `json:"items"`
}
