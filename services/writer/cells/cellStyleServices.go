package cells

import (
	"encoding/json"

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
