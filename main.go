package main

import (
	"net/http"
	"rinha-de-backend-2023/controllers"
)

func main() {

	http.HandleFunc("POST /pessoas", controllers.InsertPeople)
	http.HandleFunc("GET /pessoas/{id}", controllers.SearchByID)
	http.HandleFunc("GET /pessoas", controllers.TermSearch)
	http.HandleFunc("GET /contagem-pessoas", controllers.Count)
	
	http.ListenAndServe(":3000", nil)
}