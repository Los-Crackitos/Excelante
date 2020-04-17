package services

import (
	"encoding/json"
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/Los-Crackitos/Excelante/models"
)

// Declare alphabet to match column header position dynamically
var alphabet = []string{"A", "B", "C", "D", "E", "F"}

// WriteExcel ...
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

// WriteSheet ...
func WriteSheet(excelSheet *models.ExcelSheets, f *excelize.File) {
	f.NewSheet(excelSheet.Name)

	if excelSheet.Mode == "Array" {
		createArray(excelSheet, f)
	}
}

func createArray(excelSheet *models.ExcelSheets, f *excelize.File) {
	for n := 0; n < len(excelSheet.Headers); n++ {
		WriteHeaders(excelSheet.Name, excelSheet.Headers[n], f, n)
	}
	// write data
}

// WriteHeaders ...
func WriteHeaders(sheetName string, excelHeaders *models.ExcelHeaders, f *excelize.File, index int) {
	f.SetCellStyle(sheetName,
		fmt.Sprintf("%s1", alphabet[index]),
		fmt.Sprintf("%s1", alphabet[index]),
		GenerateCellStyleID(excelHeaders.CellStyle, f))
	f.SetCellValue(sheetName,
		 fmt.Sprintf("%s1", alphabet[index]),
		 excelHeaders.Value)
}

// GenerateCellStyleID ...
func GenerateCellStyleID(cellStyle *models.CellStyle, f *excelize.File) int {
	styleAsByte, _ := json.Marshal(cellStyle)
	style, _ := f.NewStyle(string(styleAsByte))
	return style
}
