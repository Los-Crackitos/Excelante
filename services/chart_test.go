package services

import (
	"bytes"
	"testing"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/stretchr/testify/assert"

	"github.com/Los-Crackitos/Excelante/models"
)

func TestCreateChart(t *testing.T) {
	file := createTestFile()

	seriesSlice := []*models.Series{}

	chart := &models.Chart{
		Type: "areaStacked",
		Series: append(seriesSlice, &models.Series{
			Name:   "My Test Chart",
			Values: "Test",
		}),
		Legend: &models.Legend{
			Position: "left",
		},
	}

	file.createChart("test", "B5", chart)

	var buffer bytes.Buffer

	assert.NoError(t, file.Excel.Write(&buffer))

	newFile, err := excelize.OpenReader(&buffer)
	assert.NoError(t, err)

	_, ok := newFile.XLSX["xl/drawings/drawing1.xml"]

	assert.True(t, ok)
}
