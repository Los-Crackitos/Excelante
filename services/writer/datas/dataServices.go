package datas

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/Los-Crackitos/Excelante/models"
	"github.com/Los-Crackitos/Excelante/services/writer/cells"
	"github.com/Los-Crackitos/Excelante/services/writer/headers"
)

// WriteDatas : Writes datas into excel file
// @Param : excelSheet (*models.ExcelSheets) file (*excelize.File)
func WriteDatas(excelSheet *models.ExcelSheets, file *excelize.File) {
	for n := 0; n < len(excelSheet.Datas); n++ {
		for i := 0; i < len(excelSheet.Datas[n]); i++ {
			col, err := headers.GetNextColIndex(i)
			if err != nil {
				return
			}
			line := n + 2
			file.SetCellStyle(excelSheet.Name,
				fmt.Sprintf("%s%d", col, line),
				fmt.Sprintf("%s%d", col, line),
				cells.GenerateCellStyleID(excelSheet.Datas[n][i].CellStyle, file))
			file.SetCellValue(excelSheet.Name, fmt.Sprintf("%s%d", col, line), excelSheet.Datas[n][i].Value)
		}
	}
}
