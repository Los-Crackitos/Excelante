package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/Los-Crackitos/Excelante/models"
	services "github.com/Los-Crackitos/Excelante/services/writer"
)

// WriteExcel : Http Handler
// @Param : w (http.ResponseWriter) r (*http.Request)
func WriteExcel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var excelModel *models.ExcelModel
	if err := json.NewDecoder(r.Body).Decode(&excelModel); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode("Body malformed data")

		return
	}
	file := excelize.NewFile()
	services.WriteExcel(file, excelModel)

	/*w.Header().Set("Content-Type", "application/octet-stream")
  w.Header().Set("Content-Disposition", "attachment;filename='userInputData.xlsx'")
  w.Header().Set("File-Name", "userInputData.xlsx")
  w.Header().Set("Content-Transfer-Encoding", "binary")
  w.Header().Set("Expires", "0")
  if err := file.Write(w); err != nil {
		json.NewEncoder(w).Encode(err)
	}*/
	 json.NewEncoder(w).Encode("Ok")
}
