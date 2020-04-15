package models

//ExcelHeaders : Represent excel headers columns
type ExcelHeaders struct {
<<<<<<< HEAD
	Name string
 	CellStyle *CellStyle
=======
	Value     string          `json:"value"`
	CellStyle *CellStyle      `json:"cellStyle,omitempty"`
	Headers   []*ExcelHeaders `json:"headers,omitempty"`
>>>>>>> Clean up & refacto & add json input format
}
