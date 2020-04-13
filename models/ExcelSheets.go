package models

//ExcelSheets : Represent excel sheet
type ExcelSheets struct {
	Mode string // -> Enum?
	Name string
	Headers []*ExcelHeaders
	Datas [][] string
}
