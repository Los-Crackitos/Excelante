package models

//Protection : Represent a cell style object
type Protection struct {
	Hidden bool `json:"hidden"`
	Locked bool `json:"locked"`
}
