package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Los-Crackitos/Excelante/models"
)

// WriteExcel ...
func WriteExcel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var excelModel *models.ExcelModel
	if err := json.NewDecoder(r.Body).Decode(&excelModel); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode("Body malformed data")

		return
	}

	json.NewEncoder(w).Encode(&excelModel)
}
