package services

import (
	"encoding/json"
	"errors"

	"github.com/Los-Crackitos/Excelante/models"
)

func (file *File) generateStyle(cellStyle *models.Style) (int, error) {
	styleProps, err := json.Marshal(cellStyle)
	if err != nil {
		return 0, errors.New("An error occurred when decode cellStyle JSON data")
	}

	style, err := file.Excel.NewStyle(string(styleProps))
	if err != nil {
		return 0, errors.New("An error occurred when generating cell style")
	}

	return style, nil
}

func (file *File) writeCell(sheetName string, value interface{}, cellStyle *models.Style, cellPosition string) error {
	style, err := file.generateStyle(cellStyle)

	if err != nil {
		return err
	}

	file.Excel.SetCellStyle(sheetName,
		cellPosition,
		cellPosition,
		style)
	file.Excel.SetCellValue(sheetName, cellPosition, value)

	return nil
}
