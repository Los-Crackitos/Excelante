package services

import (
	"encoding/json"

	"github.com/Los-Crackitos/Excelante/models"
)

func (file *File) createChart(sheetName string, startingCoordinates string, chart *models.Chart) {
	chartProps, _ := json.Marshal(chart)
	file.Excel.AddChart(sheetName, startingCoordinates, string(chartProps))
}
