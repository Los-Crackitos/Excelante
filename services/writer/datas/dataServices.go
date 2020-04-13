package datas

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/Los-Crackitos/Excelante/models"
	headers "github.com/Los-Crackitos/Excelante/services/writer/headers"
)

// WriteDatas : Writes datas into excel file
// @Param : excelSheet (*models.ExcelSheets) file (*excelize.File)
func WriteDatas(excelSheet *models.ExcelSheets, file *excelize.File) {
	for n := 0; n < len(excelSheet.Datas); n++ {
		for i := 0; i < len(excelSheet.Datas[n]); i++ {
			col := headers.GetNextColIndex(i)
			line := n + 2
			file.SetCellValue(excelSheet.Name, fmt.Sprintf("%s%s" ,col, line), excelSheet.Datas[n][i])
		}
	}
}
