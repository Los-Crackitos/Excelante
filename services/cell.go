package services

import (
	"encoding/json"

	"github.com/Los-Crackitos/Excelante/models"
)

func (file *File) generateStyle(cellStyle *models.Style) int {
	styleProps, _ := json.Marshal(cellStyle)
	style, _ := file.Excel.NewStyle(string(styleProps))

	return style
}

func (file *File) writeCell(sheetName string, value interface{}, cellStyle *models.Style, cellPosition string) error {
	style := file.generateStyle(cellStyle)

	file.Excel.SetCellStyle(sheetName,
		cellPosition,
		cellPosition,
		style)
	file.Excel.SetCellValue(sheetName, cellPosition, value)

	return nil
}
