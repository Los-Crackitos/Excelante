package models

//ExcelModel : Represent excel file model
type ExcelModel struct {
	Sheets []*ExcelSheets `json:"sheets"`
}
