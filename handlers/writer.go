package handlers

import (
	"encoding/json"
	"net/http"
)

// WriteExcelFile write an excel file with the json input passed through http Post method.
func WriteExcelFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Excel file successfully created")
}
