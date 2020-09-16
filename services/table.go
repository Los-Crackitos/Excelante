package services

import (
	"github.com/Los-Crackitos/Excelante/models"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func (file *File) createTable(sheetName string, startingCoordinates string, values []*models.ColOrRowValues) {
	for index, value := range values {
		// initialPositions is an integer array[col,row]
		var initialPositions []int
		columnPosition, rowPosition, err := excelize.CellNameToCoordinates(startingCoordinates)
		if err != nil {
			columnPosition, rowPosition = 1, 1 // StartingCoordonates are not valid, so table begin at position 1,1
		}

		// Depending table orientation, we iterate through column or row
		if value.Orientation == "row" {
			rowPosition = rowPosition + index
		} else {
			columnPosition = columnPosition + index
		}

		initialPositions = append(initialPositions, columnPosition, rowPosition)

		file.writeTable(sheetName, value.Cells, initialPositions, value.Orientation)
	}
}

func (file *File) writeTable(sheetName string, cells []*models.Cell, initialPositions []int, orientation string) {
	var columnPosition, rowPosition int

	for index, cell := range cells {
		if orientation == "row" {
			columnPosition = initialPositions[0] + index
			rowPosition = initialPositions[1]
		} else {
			columnPosition = initialPositions[0]
			rowPosition = initialPositions[1] + index
		}

		cellPosition, _ := excelize.CoordinatesToCellName(columnPosition, rowPosition)
		file.writeCell(sheetName, cell, cellPosition)
	}
}
