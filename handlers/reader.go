package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Los-Crackitos/Excelante/models"
	"github.com/Los-Crackitos/Excelante/services"
)

// ReadExcelFileByLine read a given excel file line by line and return every data into JSON format
// Max excel size is 60mb
// @Summary Read Excel file line by line
// @Description transform given Excel file to JSON
// @Tags Readers
// @Accept mpfd
// @Produce json
// @Param file body string true "The Excel file to convert"
// @Param options body models.ReaderOption false "Reader optional options"
// @Success 200 {object} services.Output
// @Failure 400 {string} string
// @Router /read/lines [post]
func ReadExcelFileByLine(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	file, _, _ := r.FormFile("file")
	r.ParseMultipartForm(60 << 20)
	options := r.Form.Get("options")

	readerOptions := models.ReaderOption{}
	json.Unmarshal([]byte(options), &readerOptions)

	output, err := services.ReadLines(file, readerOptions)

	if err != nil {
		http.Error(w, "An error occurred during file reading", http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

// ReadExcelFileByColumn read a given excel file column by column and return every data into JSON format
// Max excel size is 60mb
// @Summary Read Excel file column by column
// @Description transform given Excel file to JSON
// @Tags Readers
// @Accept mpfd
// @Produce json
// @Param file body string true "The Excel file to convert"
// @Param options body models.ReaderOption false "Reader optional options"
// @Success 200 {object} services.Output
// @Failure 400 {string} string
// @Router /read/columns [post]
func ReadExcelFileByColumn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	file, _, _ := r.FormFile("file")
	r.ParseMultipartForm(60 << 20)
	options := r.Form.Get("options")

	readerOptions := models.ReaderOption{}
	json.Unmarshal([]byte(options), &readerOptions)

	output, err := services.ReadColumns(file, readerOptions)

	if err != nil {
		http.Error(w, "An error occurred during file reading", http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
