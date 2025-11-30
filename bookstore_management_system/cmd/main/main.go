package main

import (
	"bookstore_management_system/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookstoreRoutes(router)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}

}