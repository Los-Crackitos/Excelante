package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Los-Crackitos/Excelante/services"
)

// ReadExcelFileByLine read a given excel file line by line and return every data into JSON format
func ReadExcelFileByLine(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	output, err := services.ReadLines(file)

	if err != nil {
		http.Error(w, "An error occurred during file reading", http.StatusBadRequest)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
