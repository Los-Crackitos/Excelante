package router

import (
	"github.com/Los-Crackitos/Excelante/handlers"
	"github.com/gorilla/mux"
)

func createWriterRouter(router *mux.Router) {
	userRouter := router.PathPrefix("/write").Subrouter()

	userRouter.
		HandleFunc("", handlers.WriteExcel).
		Methods("POST")
}
