package services

import (
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// Output is a generic type used by Excel reading functions
type Output map[string]map[int]interface{}

// ReadLines read all lines of a given Excel file
// Returns all values of the file using the Output type or an error
func ReadLines(file *excelize.File, sheet string, ch chan Output) error {
	output := make(Output)

	fmt.Println("SHEET INSIDE FUNCTION", sheet)

	for _, sheetName := range file.GetSheetMap() {
		if sheetName != sheet {
			continue
		}

		output[sheetName] = make(map[int]interface{})

		rows, err := file.Rows(sheetName)
		if err != nil {
			return err
		}

		currentRowIndex := 1

		for rows.Next() {
			row, err := rows.Columns()
			if err != nil {
				return err
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

	ch <- output

	return nil
}

// ReadColumns read all columns of a given Excel file
// Returns all values of the file using the Output type or an error
func ReadColumns(file multipart.File, sheet string, ch chan Output) error {
	output := make(Output)
	f, err := excelize.OpenReader(file)

	if err != nil {
		return err
	}

	for _, sheetName := range f.GetSheetMap() {
		if sheetName != sheet {
			continue
		}

		output[sheetName] = make(map[int]interface{})

		cols, err := f.Cols(sheetName)
		if err != nil {
			return err
		}

		currentColIndex := 1

		for cols.Next() {
			col, err := cols.Rows()
			if err != nil {
				return err
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

	ch <- output

	return nil
}
