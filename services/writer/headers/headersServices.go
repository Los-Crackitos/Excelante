package headers

import (
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
func WriteHeader(sheetName string, excelHeader *models.ExcelHeaders, file *excelize.File, columnIndex int) {
	lineIndex := 1
	cellCoordonates, err := cells.GetCellCoordonates(columnIndex, lineIndex)
	if err != nil {
		return
	}

	file.SetCellStyle(sheetName,
		cellCoordonates,
		cellCoordonates,
		cells.GenerateCellStyleID(excelHeader.CellStyle, file))
	file.SetCellValue(sheetName,
		cellCoordonates,
		excelHeader.Value)
}
