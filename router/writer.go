package router

import (
	"github.com/Los-Crackitos/Excelante/handlers"
	"github.com/gorilla/mux"
)

// CreateWriterRouter : Create writer routes
func CreateWriterRouter(router *mux.Router) {
	writerRouter := router.PathPrefix("/write").Subrouter()

	writerRouter.
		HandleFunc("", handlers.WriteExcelFile).
		Methods("POST")
}
