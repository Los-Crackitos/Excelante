package services

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/Los-Crackitos/Excelante/models"
	datas "github.com/Los-Crackitos/Excelante/services/writer/datas"
	headers "github.com/Los-Crackitos/Excelante/services/writer/headers"
)

// WriteExcel : Create and write datas into excel file
// @Param : file (*excelize.File) excelModel (*models.ExcelModel)
func WriteExcel(file *excelize.File, excelModel *models.ExcelModel) {
		for n := 0; n < len(excelModel.Sheets); n++ {
			file.SetActiveSheet(n)
			WriteSheet(excelModel.Sheets[n], file)
		}
		file.DeleteSheet("Sheet1") // By default, excelize.NewFile create a default sheet call "Sheet1", then we delete it.

    if err := file.SaveAs("Book1.xlsx"); err != nil {
        fmt.Println(err)
		}
}

// WriteSheet : Create and write into excel sheet
// @Param : excelSheet (*models.ExcelSheets) file (*excelize.File)
func WriteSheet(excelSheet *models.ExcelSheets, file *excelize.File) {
	f.NewSheet(excelSheet.Name)

	if excelSheet.Mode == "Array" {
		CreateArray(excelSheet, file)
	}
	// Mode Graph to create graph into excel ...
}

// CreateArray : Create an array into excel file with the inputs datas
// @Param : excelSheet (*models.ExcelSheets) file (*excelize.File)
func CreateArray(excelSheet *models.ExcelSheets, file *excelize.File) {
	headers.WriteHeaders(excelSheet.Name, excelSheet.Headers, file)
	datas.WriteDatas(excelSheet, file)
}
