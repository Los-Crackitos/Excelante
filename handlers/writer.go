package handlers

import (
	"encoding/json"
	"net/http"
)

// WriteExcelFile :
func WriteExcelFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Ok")
}
