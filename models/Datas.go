package models

//Datas ...
type Datas struct {
	Value     string     `json:"value"`
	CellStyle *CellStyle `json:"cellstyle,omitempty"`
}
