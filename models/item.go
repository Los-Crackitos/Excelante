package models

// Item represent  json model Item
type Item struct {
	Mode                    string        `json:"mode"`
	StartingCellCoordonates string        `json:"starting_cell_coordonates,omitempty"`
	Data                    []interface{} `json:"data"`
}
