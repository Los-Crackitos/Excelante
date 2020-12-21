package services

import (
	"os"
	"testing"

	"github.com/Los-Crackitos/Excelante/models"

	"github.com/stretchr/testify/assert"
)

func TestReadLines(t *testing.T) {
	protectedFile, _ := os.Open("../test/protected_file.xlsx")
	protectedFileOutput, protectedFileErr := ReadLines(protectedFile, models.ReaderOption{})

	assert.Nil(t, protectedFileOutput)
	assert.Error(t, protectedFileErr)

	normalFile, _ := os.Open("../test/input.xlsx")
	normalFileOutput, normalFileErr := ReadLines(normalFile, models.ReaderOption{})

	assert.NoError(t, normalFileErr)

	firstRow := normalFileOutput["Feuil1"][1]

	assert.Contains(t, firstRow, "Cell A1", "First Row should contains \"Cell A1\"")
	assert.Contains(t, firstRow, "Cell B1", "First Row should contains \"Cell B1\"")
	assert.NotContains(t, firstRow, "Cell A2", "First Row should not contains \"Cell A2\"")
	assert.NotContains(t, firstRow, "Cell B3", "First Row should not contains \"Cell B3\"")

	assert.Contains(t, normalFileOutput["Feuil1"][4], "N/A", "Fourth Row should contains \"N/A\"")

	normalFile, _ = os.Open("../test/input.xlsx")
	readerOption := models.ReaderOption{
		Options: []models.Option{
			models.Option{
				SheetName:           "Feuil2",
				StartingCoordinates: "A2",
			},
		},
	}
	normalFileOutput, normalFileErr = ReadLines(normalFile, readerOption)

	assert.NoError(t, normalFileErr)

	firstRow = normalFileOutput["Feuil2"][1]
	assert.Nil(t, firstRow)

	firstRow = normalFileOutput["Feuil2"][2]

	assert.Contains(t, firstRow, "feuil2 A2", "First Row should contains \"feuil2 A2\"")
	assert.Contains(t, firstRow, "feuil2 B2", "First Row should contains \"feuil2 B2\"")
	assert.NotContains(t, firstRow, "Cell A2", "First Row should not contains \"Cell A2\"")
	assert.NotContains(t, firstRow, "Cell B3", "First Row should not contains \"Cell B3\"")

	assert.Contains(t, normalFileOutput["Feuil2"][3], "N/A", "Third Row should contains \"N/A\"")

	normalFile, _ = os.Open("../test/input.xlsx")
	readerOption = models.ReaderOption{
		Options: []models.Option{
			models.Option{
				SheetName: "Feuil2",
			},
		},
	}
	normalFileOutput, normalFileErr = ReadLines(normalFile, readerOption)

	assert.NoError(t, normalFileErr)

	firstRow = normalFileOutput["Feuil2"][1]

	assert.Contains(t, firstRow, "feuil2 A1", "First Row should contains \"feuil2 A1\"")
	assert.Contains(t, firstRow, "feuil2 B1", "First Row should contains \"feuil2 B1\"")
	assert.NotContains(t, firstRow, "Cell A2", "First Row should not contains \"Cell A2\"")
	assert.NotContains(t, firstRow, "Cell B3", "First Row should not contains \"Cell B3\"")

	assert.Contains(t, normalFileOutput["Feuil2"][3], "N/A", "Third Row should contains \"N/A\"")
}

func TestReadColumns(t *testing.T) {
	protectedFile, _ := os.Open("../test/protected_file.xlsx")
	protectedFileOutput, protectedFileErr := ReadColumns(protectedFile, models.ReaderOption{})

	assert.Nil(t, protectedFileOutput)
	assert.Error(t, protectedFileErr)

	normalFile, _ := os.Open("../test/input.xlsx")
	normalFileOutput, normalFileErr := ReadColumns(normalFile, models.ReaderOption{})

	assert.NoError(t, normalFileErr)

	firstCol := normalFileOutput["Feuil1"][1]

	assert.Contains(t, firstCol, "Cell A1", "First Col should contains \"Cell A1\"")
	assert.Contains(t, firstCol, "Cell A2", "First Col should contains \"Cell A2\"")
	assert.NotContains(t, firstCol, "Cell B1", "First Col should not contains \"Cell B1\"")
	assert.NotContains(t, firstCol, "Cell B2", "First Col should not contains \"Cell B2\"")

	assert.Contains(t, firstCol, "N/A", "Fourth Col should contains \"N/A\"")

	normalFile, _ = os.Open("../test/input.xlsx")
	readerOption := models.ReaderOption{
		Options: []models.Option{
			models.Option{
				SheetName:           "Feuil2",
				StartingCoordinates: "B2",
			},
		},
	}

	normalFileOutput, normalFileErr = ReadColumns(normalFile, readerOption)

	assert.NoError(t, normalFileErr)

	firstCol = normalFileOutput["Feuil2"][1]
	assert.Nil(t, firstCol)

	firstCol = normalFileOutput["Feuil2"][2]

	assert.Contains(t, firstCol, "feuil2 B2", "First Col should contains \"feuil2 B2\"")
	assert.Contains(t, firstCol, "feuil2 B3", "First Col should contains \"feuil2 B3\"")
	assert.NotContains(t, firstCol, "Cell B1", "First Col should not contains \"Cell A2\"")
	assert.NotContains(t, firstCol, "Cell B2", "First Col should not contains \"Cell B3\"")

	assert.Contains(t, firstCol, "N/A", "Third Col should contains \"N/A\"")

	normalFile, _ = os.Open("../test/input.xlsx")
	readerOption = models.ReaderOption{
		Options: []models.Option{
			models.Option{
				SheetName: "Feuil2",
			},
		},
	}

	normalFileOutput, normalFileErr = ReadColumns(normalFile, readerOption)

	assert.NoError(t, normalFileErr)

	firstCol = normalFileOutput["Feuil2"][1]

	assert.Contains(t, firstCol, "feuil2 A1", "First Col should contains \"feuil2 A1\"")
	assert.Contains(t, firstCol, "feuil2 A2", "First Col should contains \"feuil2 A2\"")
	assert.NotContains(t, firstCol, "Cell B1", "First Col should not contains \"Cell A2\"")
	assert.NotContains(t, firstCol, "Cell B2", "First Col should not contains \"Cell B3\"")

	assert.Contains(t, firstCol, "N/A", "Third Col should contains \"N/A\"")
}
