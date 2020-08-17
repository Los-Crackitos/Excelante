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

	err := file.writeCell("test", "myTestValue", styles, "A2")
	assert.Nil(t, err)

	value, err := file.Excel.GetCellValue("test", "A2")
	assert.Nil(t, err)
	assert.Equal(t, value, "myTestValue", "Value should be \"myTestValue\"")

	styleIndex, err := file.Excel.GetCellStyle("test", "A2")
	assert.Nil(t, err)
	assert.Equal(t, styleIndex, 1, "Style index should be \"1\"")
}
