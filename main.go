package main

import (
	"net/http"
	"rinha-de-backend-2023/controllers"
)

func main() {

	http.HandleFunc("/pessoas", controllers.InsertPeople)
	http.HandleFunc("/pessoas/{id}", controllers.SearchByID)

	http.ListenAndServe(":3000", nil)
}