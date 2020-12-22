package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/Los-Crackitos/Excelante/services"
	"github.com/imdario/mergo"
)

// ReadExcelFileByLine read a given excel file line by line and return every data into JSON format
// @Summary Read Excel file line by line
// @Description transform given Excel file to JSON
// @Tags Readers
// @Accept mpfd
// @Produce json
// @Param file body string true "The Excel file to convert"
// @Param sheets body string false "Sheets to extract"
// @Success 200 {object} services.Output
// @Failure 400 {string} string
// @Router /read/lines [post]
func ReadExcelFileByLine(w http.ResponseWriter, r *http.Request) {
	log.Println("yo")
	file, _, _ := r.FormFile("file")
	log.Println("FILE", file)
	r.ParseForm()
	sheets := r.Form.Get("sheets")

	var sheetsToExtract []string
	if sheets != "" {
		sheetsToExtract = strings.Split(sheets, ",")
	}

	ch := make(chan services.Output)
	var wg sync.WaitGroup
	f, err := excelize.OpenReader(file)

	if err != nil {
		http.Error(w, "An error occurred during file reading", http.StatusBadRequest)

		return
	}
	output := make(services.Output)

	for _, sheet := range sheetsToExtract {
		go services.ReadLines(f, sheet, ch)
		wg.Add(1)
		go consume(output, ch, &wg)
	}
	wg.Wait()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func consume(output services.Output, ch chan services.Output, wg *sync.WaitGroup) {
	defer wg.Done()

	if err := mergo.Merge(&output, <-ch); err != nil {
		log.Println("merge err", err)
	}
}

// ReadExcelFileByColumn read a given excel file column by column and return every data into JSON format
// @Summary Read Excel file column by column
// @Description transform given Excel file to JSON
// @Tags Readers
// @Accept mpfd
// @Produce json
// @Param file body string true "The Excel file to convert"
// @Param sheets body string false "Sheets to extract"
// @Success 200 {object} services.Output
// @Failure 400 {string} string
// @Router /read/columns [post]
func ReadExcelFileByColumn(w http.ResponseWriter, r *http.Request) {
	file, _, _ := r.FormFile("file")
	r.ParseForm()
	sheets := r.Form.Get("sheets")

	var sheetsToExtract []string
	if sheets != "" {
		sheetsToExtract = strings.Split(sheets, ",")
	}

	ch := make(chan services.Output)

	for _, sheet := range sheetsToExtract {
		go services.ReadColumns(file, sheet, ch)

		// if err != nil {
		// 	http.Error(w, "An error occurred during file reading", http.StatusBadRequest)

		// 	return
		// }
	}

	log.Println(<-ch)

	var output string

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
