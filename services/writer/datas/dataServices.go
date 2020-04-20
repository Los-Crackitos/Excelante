package datas

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/Los-Crackitos/Excelante/models"
	"github.com/Los-Crackitos/Excelante/services/writer/cells"
)

// WriteDatas : Writes datas into excel file
// @Param : excelSheet (*models.ExcelSheets) file (*excelize.File)
func WriteDatas(excelSheet *models.ExcelSheets, file *excelize.File) {
	for n := 0; n < len(excelSheet.Datas); n++ {
		for i := 0; i < len(excelSheet.Datas[n]); i++ {
			columnIndex := i + 1
			line := n + 2
			cellCoordonates, err := cells.GetCellCoordonates(columnIndex, line)
			if err != nil {
				return
			}
			file.SetCellStyle(excelSheet.Name,
				cellCoordonates,
				cellCoordonates,
				cells.GenerateCellStyleID(excelSheet.Datas[n][i].CellStyle, file))
			file.SetCellValue(excelSheet.Name, cellCoordonates, excelSheet.Datas[n][i].Value)
		}
	}
}
