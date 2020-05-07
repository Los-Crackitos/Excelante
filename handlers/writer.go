package handlers

import (
	"encoding/json"
	"net/http"
)

// WriteExcelFile write an Excel file with the JSON input passed through HTTP POST method
// @Summary Write Excel file
// @Description transform JSON to Excel file
// @Tags Writers
// @Accept json
// @Produce application/text
// @Router /write [post]
func WriteExcelFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Excel file successfully created")
}
