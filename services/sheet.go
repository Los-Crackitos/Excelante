package services

import (
	"github.com/Los-Crackitos/Excelante/models"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// File contain the excelize file type and is used as service methods receiver
type File struct {
	Excel *excelize.File
}

// WriteSheets write a sheets into a given Excel file
func (file *File) WriteSheets(sheets []*models.Sheet) {
	for _, sheet := range sheets {
		file.writeSheet(sheet)
	}
	file.Excel.DeleteSheet("Sheet1")
}

func (file *File) writeSheet(sheet *models.Sheet) {
	file.Excel.NewSheet(sheet.Name)
	file.Excel.SetPageLayout(sheet.Name, excelize.PageLayoutOrientation(sheet.Orientation))
	file.writeItems(sheet.Name, sheet.Items)
}
