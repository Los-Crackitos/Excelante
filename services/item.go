package services

import "github.com/Los-Crackitos/Excelante/models"

func (file *File) writeItems(sheetName string, items []*models.Item) {
	for _, item := range items {
		file.writeItem(sheetName, item)
	}
}

func (file *File) writeItem(sheetName string, item *models.Item) {
	if item.Chart != nil {
		file.createChart(sheetName, item.StartingCellCoordinates, item.Chart)
		return
	}
	file.createTable(sheetName, item.StartingCellCoordinates, item.Table)
}
