package services

import (
	"encoding/json"
	"errors"

	"github.com/Los-Crackitos/Excelante/models"
)

func (file *File) createGraph(sheetName string, startingCoordinates string, graph *models.Graph) error {
	graphProps, err := json.Marshal(graph)
	if err != nil {
		return errors.New("An error occurred when decode graph JSON data")
	}
	file.Excel.AddChart(sheetName, startingCoordinates, string(graphProps))
	return nil
}
