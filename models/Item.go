package models

// Item : Represent json model Item
type Item struct {
	Mode                    string  `json:"mode"`
	StartingCellCoordonates string  `json:"startingCellCoordonates,omitempty"`
	Datas                   []*Data `json:"datas"`
}
