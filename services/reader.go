package services

import (
	"mime/multipart"
	"strings"

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
func ReadLines(file multipart.File) (Output, error) {
	output := make(Output)
	f, err := excelize.OpenReader(file)

	if err != nil {
		return nil, err
	}

	for _, sheetName := range f.GetSheetMap() {
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

			rowValue := make([]interface{}, 0)

			for _, cellValue := range row {
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
func ReadColumns(file multipart.File) (Output, error) {
	output := make(Output)
	f, err := excelize.OpenReader(file)

	if err != nil {
		return nil, err
	}

	for _, sheetName := range f.GetSheetMap() {
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

			rowValue := make([]interface{}, 0)

			for _, cellValue := range col {
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
