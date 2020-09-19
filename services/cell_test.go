package services

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Los-Crackitos/Excelante/models"
)

func TestGenerateStyle(t *testing.T) {
	file := createTestFile()

	borderSlice := []*models.Border{}
	styles := &models.Style{
		Border: append(borderSlice, &models.Border{
			Type:  "top",
			Color: "#000000",
			Style: 1,
		}),
	}

	style := file.generateStyle(styles)

	assert.Equal(t, style, 1, "Style should be 1")
}

func TestWriteCell(t *testing.T) {
	file := createTestFile()

	borderSlice := []*models.Border{}
	styles := &models.Style{
		Border: append(borderSlice, &models.Border{
			Type:  "top",
			Color: "#000000",
			Style: 1,
		}),
	}

	cell := &models.Cell{
		Value: "myTestValue",
		Style: styles,
	}

	err := file.writeCell("test", cell, "A2")
	assert.Nil(t, err)

	value, err := file.Excel.GetCellValue("test", "A2")
	assert.Nil(t, err)
	assert.Equal(t, value, "myTestValue", "Value should be \"myTestValue\"")

	styleIndex, err := file.Excel.GetCellStyle("test", "A2")
	assert.Nil(t, err)
	assert.Equal(t, styleIndex, 1, "Style index should be \"1\"")

	cellWithFormula := &models.Cell{
		Value:     "SUM(B2:B5)",
		Style:     styles,
		IsFormula: true,
	}

	err = file.writeCell("test", cellWithFormula, "A5")
	assert.Nil(t, err)

	value, err = file.Excel.GetCellFormula("test", "A5")
	assert.Nil(t, err)
	assert.Equal(t, value, "SUM(B2:B5)", "Value should be \"SUM(B2:B5)\"")
}
