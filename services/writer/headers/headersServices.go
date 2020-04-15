package headers

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/Los-Crackitos/Excelante/models"
	"github.com/Los-Crackitos/Excelante/services/writer/cells"
)

// WriteHeaders ...
// @Param :  sheetName (string) excelHeader (*models.ExcelHeaders) file (*excelize.File)
func WriteHeaders(sheetName string, excelHeaders []*models.ExcelHeaders, f *excelize.File) {
	for n := 0; n < len(excelHeaders); n++ {
		WriteHeader(sheetName, excelHeaders[n], f, n)
	}
}

// WriteHeader : Write header into excel file at the specified index
// @Param : sheetName (string) excelHeader (*models.ExcelHeaders) file (*excelize.File), index (int)
func WriteHeader(sheetName string, excelHeader *models.ExcelHeaders, file *excelize.File, index int) {
	colIndex, err := GetNextColIndex(index)
	if err != nil {
		return
	}

	file.SetCellStyle(sheetName,
		fmt.Sprintf("%s1", colIndex),
		fmt.Sprintf("%s1", colIndex),
		cells.GenerateCellStyleID(excelHeader.CellStyle, file))
	file.SetCellValue(sheetName,
		fmt.Sprintf("%s1", colIndex),
		excelHeader.Value)
}

// GetNextColIndex : Return column index (A, B, C, ...) depending parent loop index
// @Param : index (int)
// @Return : columnIndex (string) error (error)
func GetNextColIndex(index int) (string, error) {
	if index < 1 {
		return "", fmt.Errorf("index (%d) must be positive and superior than 1", index)
	}
	var col string
	for index > 0 {
		col = string((index-1)%26+65) + col
		index = (index - 1) / 26
	}
	return col, nil
}
