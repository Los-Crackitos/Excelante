package services

import (
	"bytes"
	"testing"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/stretchr/testify/assert"

	"github.com/Los-Crackitos/Excelante/models"
)

func TestWriteItems(t *testing.T) {
	file := createTestFile()

	file.writeItems("test", append(items, &models.Item{
		StartingCellCoordinates: "D4",
		Table: append(tables, &models.ColOrRowValues{
			Orientation: "row",
			Cells: append(cells, &models.Cell{
				Value: "Test Cell Row 1",
			}, &models.Cell{
				Value: "Test Cell Row 2",
			}),
		}),
	}, &models.Item{
		StartingCellCoordinates: "I10",
		Table: append(tables, &models.ColOrRowValues{
			Orientation: "col",
			Cells: append(cells, &models.Cell{
				Value: "Test Cell Col 1",
			}, &models.Cell{
				Value: "Test Cell Col 2",
			}),
		}),
	}, &models.Item{
		StartingCellCoordinates: "G",
		Table: append(tables, &models.ColOrRowValues{
			Cells: append(cells, &models.Cell{
				Value: "Test Cell 1",
			}),
		}),
	}))

	firstRowCellValue, err := file.Excel.GetCellValue("test", "D4")
	assert.NoError(t, err)
	assert.Equal(t, firstRowCellValue, "Test Cell Row 1", "Cell value should be \"Test Cell Row 1\"")

	secondRowCellValue, err := file.Excel.GetCellValue("test", "E4")
	assert.NoError(t, err)
	assert.Equal(t, secondRowCellValue, "Test Cell Row 2", "Cell value should be \"Test Cell Row 2\"")

	UnknownCellValue, err := file.Excel.GetCellValue("test", "D5")
	assert.NoError(t, err)
	assert.Empty(t, UnknownCellValue, "Cell value should be empty")

	firstColCellValue, err := file.Excel.GetCellValue("test", "I10")
	assert.NoError(t, err)
	assert.Equal(t, firstColCellValue, "Test Cell Col 1", "Cell value should be \"Test Cell Col 1\"")

	secondColCellValue, err := file.Excel.GetCellValue("test", "I11")
	assert.NoError(t, err)
	assert.Equal(t, secondColCellValue, "Test Cell Col 2", "Cell value should be \"Test Cell Col 2\"")

	testCellValue, err := file.Excel.GetCellValue("test", "A1")
	assert.NoError(t, err)
	assert.Equal(t, testCellValue, "Test Cell 1", "Cell value should be \"Test Cell 1\"")
}
func TestWriteRowItem(t *testing.T) {
	file := createTestFile()

	file.writeItem("test", &models.Item{
		StartingCellCoordinates: "H8",
		Table: append(tables, &models.ColOrRowValues{
			Orientation: "row",
			Cells: append(cells, &models.Cell{
				Value: "Test Cell H8",
			}, &models.Cell{
				Value: "Test Cell I8",
			}),
		}),
	})

	firstCellValue, err := file.Excel.GetCellValue("test", "H8")
	assert.NoError(t, err)
	assert.Equal(t, firstCellValue, "Test Cell H8", "Cell value should be \"Test Cell H8\"")

	secondCellValue, err := file.Excel.GetCellValue("test", "I8")
	assert.NoError(t, err)
	assert.Equal(t, secondCellValue, "Test Cell I8", "Cell value should be \"Test Cell I8\"")
}

func TestWriteColItem(t *testing.T) {
	file := createTestFile()

	file.writeItem("test", &models.Item{
		StartingCellCoordinates: "H8",
		Table: append(tables, &models.ColOrRowValues{
			Orientation: "col",
			Cells: append(cells, &models.Cell{
				Value: "Test Cell H8",
			}, &models.Cell{
				Value: "Test Cell H9",
			}),
		}),
	})

	firstCellValue, err := file.Excel.GetCellValue("test", "H8")
	assert.NoError(t, err)
	assert.Equal(t, firstCellValue, "Test Cell H8", "Cell value should be \"Test Cell H8\"")

	secondCellValue, err := file.Excel.GetCellValue("test", "H9")
	assert.NoError(t, err)
	assert.Equal(t, secondCellValue, "Test Cell H9", "Cell value should be \"Test Cell H9\"")
}

func TestWriteChartItem(t *testing.T) {
	file := createTestFile()

	seriesSlice := []*models.Series{}

	file.writeItem("test", &models.Item{
		StartingCellCoordinates: "H8",
		Table: append(tables, &models.ColOrRowValues{
			Orientation: "row",
			Cells: append(cells, &models.Cell{
				Value: "Test Cell H8",
			}),
		}),
		Chart: &models.Chart{
			Type: "areaStacked",
			Series: append(seriesSlice, &models.Series{
				Name:   "My Test Chart",
				Values: "Test",
			}),
			Legend: &models.Legend{
				Position: "left",
			},
		},
	})

	var buffer bytes.Buffer

	assert.NoError(t, file.Excel.Write(&buffer))

	newFile, err := excelize.OpenReader(&buffer)
	assert.NoError(t, err)

	_, ok := newFile.XLSX["xl/drawings/drawing1.xml"]

	assert.True(t, ok)
}
