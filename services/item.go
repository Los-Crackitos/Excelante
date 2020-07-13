package services

import "github.com/Los-Crackitos/Excelante/models"

func (file *File) writeItems(sheetName string, items []*models.Item) {
	for _, item := range items {
		file.writeItem(sheetName, item)
	}
}

func (file *File) writeItem(sheetName string, item *models.Item) {
	file.createTables(sheetName, item.StartingCellCoordinates, item.Tables)
	if item.Graph != nil {
		file.createGraph(sheetName, item.StartingGraphCoordinates, item.Graph)
	}
}
