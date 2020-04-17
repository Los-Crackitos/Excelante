package services

import (
	"github.com/Los-Crackitos/Excelante/models"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func (file *File) createTables(sheetName string, startingCoordinates string, tables []*models.Table) {
	for index, table := range tables {
		file.writeTable(sheetName, table.Cells, index+1, table.Orientation)
	}
}

func (file *File) writeTable(sheetName string, cells []*models.Cell, initialPosition int, orientation string) {
	var columnPosition, rowPosition int

	for index, cell := range cells {
		if orientation == "row" {
			columnPosition = index + 1
			rowPosition = initialPosition
		} else {
			columnPosition = initialPosition
			rowPosition = index + 1
		}

		cellPosition, _ := excelize.CoordinatesToCellName(columnPosition, rowPosition)
		file.writeCell(sheetName, cell.Value, cell.Style, cellPosition)
	}
}
