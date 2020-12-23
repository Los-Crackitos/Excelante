package models

// ReaderOption contains reader optional options
type ReaderOption struct {
	Options []Option `json:"options"`
}

// Option represent a reader option object
type Option struct {
	SheetName           string `json:"sheet_name"`
	StartingCoordinates string `json:"starting_coordinates,omitempty"`
}
