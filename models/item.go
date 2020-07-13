package models

// Item struct is a thing (table, graph, ...) you want to insert into your Excel file
type Item struct {
	Mode                     string   `json:"mode"`
	StartingCellCoordinates  string   `json:"starting_cell_coordinates,omitempty"`
	StartingGraphCoordinates string   `json:"starting_graph_coordinates,omitempty"`
	Tables                   []*Table `json:"tables"`
	Graph                    *Graph   `json:"graph,omitempty"`
}
