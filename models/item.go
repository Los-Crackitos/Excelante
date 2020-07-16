package models

// Item struct is a thing (table, graph, ...) you want to insert into your Excel file
type Item struct {
	StartingCellCoordinates string   `json:"starting_cell_coordinates,omitempty"`
	Tables                  []*Table `json:"tables"`
	Graph                   *Graph   `json:"graph,omitempty"`
}
