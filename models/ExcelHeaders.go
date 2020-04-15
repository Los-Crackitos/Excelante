package models

//ExcelHeaders : Represent excel headers columns
type ExcelHeaders struct {
	Value     string          `json:"value"`
	CellStyle *CellStyle      `json:"cellStyle,omitempty"`
	Headers   []*ExcelHeaders `json:"headers,omitempty"`
}
