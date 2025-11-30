package routes

import (
	"bookstore_management_system/pkg/controllers"

	"github.com/gorilla/mux"
)
func RegisterBookstoreRoutes(router *mux.Router){
	router.HandleFunc("/book", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
}
