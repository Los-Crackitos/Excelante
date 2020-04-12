package models

// ExcelSheets ...
type ExcelSheets struct {
	Mode string // -> Enum?
	Name string
	Headers []*ExcelHeaders
}
