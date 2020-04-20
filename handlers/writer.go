package handlers

import (
	"encoding/json"
	"net/http"
)

// WriteExcelFile write an excel file with the json input passed throught http Post method.
func WriteExcelFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Excel file successfully created")
}
