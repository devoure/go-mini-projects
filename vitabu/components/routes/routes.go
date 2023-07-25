package routes

import (
	"vitabu/components/controllers"

	"github.com/gorilla/mux"
)

func RegisterSiteRoutes(router *mux.Router) {
	router.HandleFunc("/vitabu", controllers.GetAllBooksController).Methods("GET")
	router.HandleFunc("/vitabu/{book_id}", controllers.GetBookByIdController).Methods("GET")
	router.HandleFunc("/vitabu/create", controllers.CreateBookRecordController).Methods("POST")
	router.HandleFunc("/vitabu/delete/{book_id}", controllers.DeleteBookRecordController).Methods("DELETE")
	router.HandleFunc("/vitabu/update/{book_id}", controllers.UpdateBookRecordController).Methods("PUT")
}
