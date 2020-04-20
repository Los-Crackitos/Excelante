package router

import (
	"github.com/Los-Crackitos/Excelante/handlers"
	"github.com/gorilla/mux"
)

// CreateWriterRouter create writer routes and dispatch http method post to handler
func createWriterRouter(router *mux.Router) {
	writerRouter := router.PathPrefix("/write").Subrouter()

	writerRouter.
		HandleFunc("", handlers.WriteExcelFile).
		Methods("POST")
}
