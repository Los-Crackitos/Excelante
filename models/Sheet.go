package models

// Sheet : Represent json model Sheet
type Sheet struct {
	Name        string  `json:"name"`
	Orientation string  `json:"orientation,omitempty"` // Défault : landscape
	Items       []*Item `json:"items"`
}
