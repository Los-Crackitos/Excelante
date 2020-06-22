package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Los-Crackitos/Excelante/services"
)

// ReadExcelFileByLine read a given excel file line by line and return every data into JSON format
// @Summary Read Excel file line by line
// @Description transform given Excel file to JSON
// @Tags Readers
// @Accept mpfd
// @Produce json
// @Param file body string true "The Excel file to convert"
// @Success 200 {object} services.Output
// @Failure 400 {string} string
// @Router /read/lines [post]
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

// ReadExcelFileByColumn read a given excel file column by column and return every data into JSON format
// @Summary Read Excel file column by column
// @Description transform given Excel file to JSON
// @Tags Readers
// @Accept mpfd
// @Produce json
// @Param file body string true "The Excel file to convert"
// @Success 200 {object} services.Output
// @Failure 400 {string} string
// @Router /read/columns [post]
func ReadExcelFileByColumn(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	output, err := services.ReadColumns(file)

	if err != nil {
		http.Error(w, "An error occurred during file reading", http.StatusBadRequest)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
