package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Los-Crackitos/Excelante/models"
	"github.com/Los-Crackitos/Excelante/services"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// WriteExcelFile write an Excel file with the JSON input passed through HTTP POST method
// @Summary Write Excel file
// @Description transform JSON to Excel file
// @Tags Writers
// @Accept json
// @Param file body models.File true "Items you want to convert into an Excel file"
// @Produce octet-stream
// @Success 200 {string} string
// @Failure 422 {string} string
// @Router /write [post]
func WriteExcelFile(w http.ResponseWriter, r *http.Request) {
	var body models.File

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Malformed body", http.StatusUnprocessableEntity)

		return
	}

	file := services.File{
		Excel: excelize.NewFile(),
	}

	file.WriteSheets(body.Sheets)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set(
		"Content-Disposition", fmt.Sprintf(
			"attachment; filename=file_generated_%s.xlsx",
			time.Now().Format("2006-01-02_15:04:05"),
		),
	)
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")

	file.Excel.Write(w)
}
