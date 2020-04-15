package models

//ExcelSheets : Represent excel sheet
type ExcelSheets struct {
	Mode    string          `json:"mode"`
	Name    string          `json:"name"`
	Headers []*ExcelHeaders `json:"headers"`
	Datas   [][]*Datas      `json:"datas"`
}
