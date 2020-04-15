package services

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/Los-Crackitos/Excelante/models"
	"github.com/Los-Crackitos/Excelante/services/writer/datas"
	"github.com/Los-Crackitos/Excelante/services/writer/headers"
)

// WriteExcel : Create and write datas into excel file
// @Param : file (*excelize.File) excelModel (*models.ExcelModel)
func WriteExcel(file *excelize.File, excelModel *models.ExcelModel) {
	for n := 0; n < len(excelModel.Sheets); n++ {
		WriteSheet(excelModel.Sheets[n], file)
	}
	file.DeleteSheet("Sheet1") // By default, excelize.NewFile create a default sheet call "Sheet1", then we delete it.

	// To delete
	if err := file.SaveAs("Book1.xlsx"); err != nil {
		fmt.Print(err)
	}
}

// WriteSheet : Create and write into excel sheet
// @Param : excelSheet (*models.ExcelSheets) file (*excelize.File)
func WriteSheet(excelSheet *models.ExcelSheets, file *excelize.File) {
	file.NewSheet(excelSheet.Name)

	if excelSheet.Mode == "table" {
		CreateTable(excelSheet, file)
	} else if excelSheet.Mode == "graph" {
		fmt.Print("graph mod")
	} else {
		return
	}
}

// CreateTable : Create an array into excel file with the inputs datas
// @Param : excelSheet (*models.ExcelSheets) file (*excelize.File)
func CreateTable(excelSheet *models.ExcelSheets, file *excelize.File) {
	headers.WriteHeaders(excelSheet.Name, excelSheet.Headers, file)
	datas.WriteDatas(excelSheet, file)
}
