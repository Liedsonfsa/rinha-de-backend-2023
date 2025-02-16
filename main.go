package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rinha-de-backend-2023/database"
	"rinha-de-backend-2023/models"
	"rinha-de-backend-2023/repositories"

	"github.com/google/uuid"
)

func main() {
	http.HandleFunc("/pessoas", postPessoas)

	http.ListenAndServe(":3000", nil)
}

func postPessoas(w http.ResponseWriter, r* http.Request) {
	var pessoa models.People
	if err := json.NewDecoder(r.Body).Decode(&pessoa); err != nil {
		fmt.Println(err)
		return
	}

	pessoa.UUID = uuid.New().String()
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}

	repositorio := repositories.NewPeopleRepository(db)
	err = repositorio.Insert(pessoa)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(pessoa); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	showPeoples()
}

func showPeoples() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repositorio := repositories.NewPeopleRepository(db)
	peoples, err := repositorio.Search()
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(peoples)
}