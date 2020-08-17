package services

import (
	"testing"

	"github.com/Los-Crackitos/Excelante/models"
	"github.com/stretchr/testify/assert"
)

func TestWriteSheets(t *testing.T) {
	file := createTestFile()

	sheets := []*models.Sheet{}

	file.WriteSheets(append(sheets, &models.Sheet{
		Name: "My Super First Sheet",
		Items: append(items, &models.Item{
			StartingCellCoordinates: "A1",
			Table: append(tables, &models.ColOrRowValues{
				Cells: append(cells, &models.Cell{
					Value: "Super First Cell",
				}),
			}),
		}),
	}, &models.Sheet{
		Name: "My Super Second Sheet",
		Items: append(items, &models.Item{
			StartingCellCoordinates: "B2",
			Table: append(tables, &models.ColOrRowValues{
				Cells: append(cells, &models.Cell{
					Value: "Super Second Cell",
				}),
			}),
		}),
	}))

	sheetList := file.Excel.GetSheetList()

	assert.Contains(t, sheetList, "My Super First Sheet")
	assert.Contains(t, sheetList, "My Super Second Sheet")
}

func TestWriteSheet(t *testing.T) {
	file := createTestFile()

	file.writeSheet(&models.Sheet{
		Name: "My Super Solo Sheet",
		Items: append(items, &models.Item{
			StartingCellCoordinates: "C3",
			Table: append(tables, &models.ColOrRowValues{
				Cells: append(cells, &models.Cell{
					Value: "Super Solo Cell",
				}),
			}),
		}),
	})

	assert.Contains(t, file.Excel.GetSheetList(), "My Super Solo Sheet")
}
