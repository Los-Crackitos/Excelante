package services

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadLines(t *testing.T) {
	protectedFile, _ := os.Open("../test/protected_file.xlsx")
	protectedFileOutput, protectedFileErr := ReadLines(protectedFile)

	assert.Nil(t, protectedFileOutput)
	assert.Error(t, protectedFileErr)

	normalFile, _ := os.Open("../test/input.xlsx")
	normalFileOutput, normalFileErr := ReadLines(normalFile)

	assert.NoError(t, normalFileErr)

	firstRow := normalFileOutput["Feuil1"][1]

	assert.Contains(t, firstRow, "Cell A1", "First Row should contains \"Cell A1\"")
	assert.Contains(t, firstRow, "Cell B1", "First Row should contains \"Cell B1\"")
	assert.NotContains(t, firstRow, "Cell A2", "First Row should not contains \"Cell A2\"")
	assert.NotContains(t, firstRow, "Cell B3", "First Row should not contains \"Cell B3\"")

	assert.Contains(t, normalFileOutput["Feuil1"][4], "N/A", "Fourth Row should contains \"N/A\"")
}

func TestReadColumns(t *testing.T) {
	protectedFile, _ := os.Open("../test/protected_file.xlsx")
	protectedFileOutput, protectedFileErr := ReadColumns(protectedFile)

	assert.Nil(t, protectedFileOutput)
	assert.Error(t, protectedFileErr)

	normalFile, _ := os.Open("../test/input.xlsx")
	normalFileOutput, normalFileErr := ReadColumns(normalFile)

	assert.NoError(t, normalFileErr)

	firstCol := normalFileOutput["Feuil1"][1]

	assert.Contains(t, firstCol, "Cell A1", "First Col should contains \"Cell A1\"")
	assert.Contains(t, firstCol, "Cell A2", "First Col should contains \"Cell A2\"")
	assert.NotContains(t, firstCol, "Cell B1", "First Col should not contains \"Cell B1\"")
	assert.NotContains(t, firstCol, "Cell B2", "First Col should not contains \"Cell B2\"")

	assert.Contains(t, firstCol, "N/A", "Fourth Col should contains \"N/A\"")

}
