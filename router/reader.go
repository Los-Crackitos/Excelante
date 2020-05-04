package router

import (
	"github.com/Los-Crackitos/Excelante/handlers"

	"github.com/gorilla/mux"
)

func createReaderRouter(router *mux.Router) {
	writerRouter := router.PathPrefix("/read").Subrouter()

	writerRouter.
		HandleFunc("/lines", handlers.ReadExcelFileByLine).
		Methods("POST")
}
