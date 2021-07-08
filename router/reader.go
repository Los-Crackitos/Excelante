package router

import (
	"github.com/Los-Crackitos/Excelante/handlers"

	"github.com/gorilla/mux"
)

func createReaderRouter(router *mux.Router) {
	readerRouter := router.PathPrefix("/read").Subrouter()

	readerRouter.
		HandleFunc("/lines", handlers.ReadExcelFileByLine).
		Methods("POST")

	readerRouter.
		HandleFunc("/columns", handlers.ReadExcelFileByColumn).
		Methods("POST")

	readerRouter.
		HandleFunc("/custom", handlers.ReadOptions).
		Methods("POST")
}
