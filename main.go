package main

import (
	"net/http"
	"rinha-de-backend-2023/controllers"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/pessoas", controllers.InsertPeople).Methods("POST")
	router.HandleFunc("/pessoas/{id}",controllers.SearchByID).Methods("GET")
	router.HandleFunc("/pessoas", controllers.TermSearch).Methods("GET")
	router.HandleFunc("/contagem-pessoas", controllers.Count).Methods("GET")
	
	http.ListenAndServe(":3000", router)
}