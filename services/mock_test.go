package services

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"

	"github.com/Los-Crackitos/Excelante/models"
)

var items []*models.Item
var tables []*models.ColOrRowValues
var cells []*models.Cell

func createTestFile() File {
	file := File{
		Excel: excelize.NewFile(),
	}

	file.Excel.NewSheet("test")

	return file
}
