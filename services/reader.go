package services

import (
	"mime/multipart"
	"strings"

	"github.com/Los-Crackitos/Excelante/models"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// Output is a generic type used by Excel reading functions
// For example:
//
// map[
//   Feuil1:
//     map[
//		   %!s(int=1): [Cell A1 Cell B1]
//		   %!s(int=2): [Cell A2 Cell B2]
//		   %!s(int=3): [Cell A3 Cell B3]
//		   %!s(int=4): [N/A Cell B4]
//			 %!s(int=5): [Cell A5 Cell B5]
//     ]
// ]
type Output map[string]map[int]interface{}

// ReadLines read all lines of a given Excel file
// Returns all values of the file using the Output type or an error
func ReadLines(file multipart.File, readerOptions models.ReaderOption) (Output, error) {
	output := make(Output)
	f, err := excelize.OpenReader(file)

	if err != nil {
		return nil, err
	}

	initialColIndex := 1
	initialRowIndex := 1
	sheetFound := false

	for _, sheetName := range f.GetSheetMap() {
		if len(readerOptions.Options) > 0 {
			sheetFound, initialColIndex, initialRowIndex = sheetFinder(sheetName, readerOptions)
			if !sheetFound {
				continue
			}
		}

		output[sheetName] = make(map[int]interface{})

		rows, err := f.Rows(sheetName)
		if err != nil {
			return nil, err
		}

		currentRowIndex := 1

		for rows.Next() {
			row, err := rows.Columns()
			if err != nil {
				return nil, err
			}

			if currentRowIndex < initialRowIndex {
				currentRowIndex++
				continue
			}

			rowValue := make([]interface{}, 0)

			for i := (initialColIndex - 1); i < len(row); i++ {
				cellValue := row[i]

				if cellValue == "" {
					cellValue = "N/A"
				}

				rowValue = append(rowValue, strings.TrimLeft(cellValue, " "))
			}

			output[sheetName][currentRowIndex] = rowValue
			currentRowIndex++
		}
	}

	return output, nil
}

// ReadColumns read all columns of a given Excel file
// Returns all values of the file using the Output type or an error
func ReadColumns(file multipart.File, readerOptions models.ReaderOption) (Output, error) {
	output := make(Output)
	f, err := excelize.OpenReader(file)

	if err != nil {
		return nil, err
	}

	initialColIndex := 1
	initialRowIndex := 1
	sheetFound := false

	for _, sheetName := range f.GetSheetMap() {
		if len(readerOptions.Options) > 0 {
			sheetFound, initialColIndex, initialRowIndex = sheetFinder(sheetName, readerOptions)
			if !sheetFound {
				continue
			}
		}

		output[sheetName] = make(map[int]interface{})

		cols, err := f.Cols(sheetName)
		if err != nil {
			return nil, err
		}

		currentColIndex := 1

		for cols.Next() {
			col, err := cols.Rows()
			if err != nil {
				return nil, err
			}

			if currentColIndex < initialColIndex {
				currentColIndex++
				continue
			}

			rowValue := make([]interface{}, 0)

			for i := (initialRowIndex - 1); i < len(col); i++ {
				cellValue := col[i]

				if cellValue == "" {
					cellValue = "N/A"
				}

				rowValue = append(rowValue, strings.TrimLeft(cellValue, " "))
			}

			output[sheetName][currentColIndex] = rowValue
			currentColIndex++
		}
	}

	return output, nil
}

func sheetFinder(sheetName string, readerOptions models.ReaderOption) (bool, int, int) {
	for _, option := range readerOptions.Options {
		if option.SheetName == sheetName {
			if option.StartingCoordinates != "" {
				initialColIndex, initialRowIndex, _ := excelize.CellNameToCoordinates(option.StartingCoordinates)
				return true, initialColIndex, initialRowIndex
			}
			return true, 1, 1
		}
	}
	return false, 1, 1
}
