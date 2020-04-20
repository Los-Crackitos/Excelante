package cells

import (
	"encoding/json"
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/Los-Crackitos/Excelante/models"
)

// GenerateCellStyleID : Convert cellstyle json string to styleID
// @Param : cellStyle (*models.CellStyle) file (*excelize.File)
// @Return : styleID (int)
func GenerateCellStyleID(cellStyle *models.CellStyle, file *excelize.File) int {
	styleAsByte, _ := json.Marshal(cellStyle)
	style, _ := file.NewStyle(string(styleAsByte))
	return style
}

// GetCellCoordonates : Return column index (A, B, C, ...) depending parent loop index
// @Param : index (int)
// @Return : columnIndex (string) error (error)
func GetCellCoordonates(columnIndex int, lineIndex int) (string, error) {
	if columnIndex < 1 {
		return "", fmt.Errorf("index (%d) must be positive and superior than 1", columnIndex)
	}
	var col string
	for columnIndex > 0 {
		col = string((columnIndex-1)%26+65) + col
		columnIndex = (columnIndex - 1) / 26
	}
	return fmt.Sprintf("%s%d", col, lineIndex), nil
}
