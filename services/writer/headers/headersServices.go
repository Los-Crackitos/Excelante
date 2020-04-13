package headers

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/Los-Crackitos/Excelante/models"
	cells "github.com/Los-Crackitos/Excelante/services/writer/cells"
)

// Declare alphabet to match column header position dynamically
var columnsIndex = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K",
	"L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

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
	colIndex := GetNextColIndex(index)

	file.SetCellStyle(sheetName,
		fmt.Sprintf("%s1" ,colIndex),
		fmt.Sprintf("%s1" ,colIndex),
		cells.GenerateCellStyleID(excelHeader.Options.CellStyle, file))
	file.SetCellValue(sheetName,
		fmt.Sprintf("%s1" ,colIndex),
		excelHeader.Name)
}

// GetNextColIndex : Return column index (A, B, C, ...) depending parent loop index
// @Param : index (int)
// @Return : columnIndex (string)
func GetNextColIndex(index int) string {
	if index >= len(columnsIndex) { // If there are more than 25 headers, header index seems to be AA, AB, AC, AD...
		index = index - len(columnsIndex)
		return fmt.Sprintf("A%s", columnsIndex[index])
	}
	return columnsIndex[index]
}
