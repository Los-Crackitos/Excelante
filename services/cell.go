package services

import (
	"encoding/json"
	"fmt"

	"github.com/Los-Crackitos/Excelante/models"
)

func (file *File) generateStyle(cellStyle *models.Style) int {
	styleProps, _ := json.Marshal(cellStyle)
	style, _ := file.Excel.NewStyle(string(styleProps))

	return style
}

func (file *File) writeCell(sheetName string, cell *models.Cell, cellPosition string) error {
	style := file.generateStyle(cell.Style)

	file.Excel.SetCellStyle(sheetName,
		cellPosition,
		cellPosition,
		style)

	if cell.IsFormula {
		file.Excel.SetCellFormula(sheetName, cellPosition, fmt.Sprintf("%v", cell.Value))
	} else {
		file.Excel.SetCellValue(sheetName, cellPosition, cell.Value)
	}

	return nil
}
