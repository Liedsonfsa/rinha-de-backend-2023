package controllers

import (
	"encoding/json"
	"net/http"
	"rinha-de-backend-2023/database"
	"rinha-de-backend-2023/models"
	"rinha-de-backend-2023/repositories"

	"github.com/google/uuid"
	// "github.com/gorilla/mux"
)

func InsertPeople(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "método não implementado", http.StatusMethodNotAllowed)
		return
	}

	var pessoa models.People
	if err := json.NewDecoder(r.Body).Decode(&pessoa); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pessoa.UUID = uuid.New().String()
	db, err := database.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	repositorio := repositories.NewPeopleRepository(db)
	err = repositorio.Insert(pessoa)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(pessoa); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func SearchByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "método não implementado", http.StatusMethodNotAllowed)
		return
	}
	
	uuid := r.URL.Path[9:]

	db, err := database.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	repository := repositories.NewPeopleRepository(db)
	people, err := repository.SearchByID(uuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(people); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Count(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "método não implementado", http.StatusMethodNotAllowed)
		return		
	}

	db, err := database.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	repository := repositories.NewPeopleRepository(db)
	quantidade, err := repository.Count()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]int{
		"quantidade": quantidade,
	}
	
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}