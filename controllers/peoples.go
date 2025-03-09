package controllers

import (
	"encoding/json"
	"net/http"
	"rinha-de-backend-2023/database"
	"rinha-de-backend-2023/models"
	"rinha-de-backend-2023/repositories"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func InsertPeople(w http.ResponseWriter, r *http.Request) {
	var people models.People
	
	if err := json.NewDecoder(r.Body).Decode(&people); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	people.UUID = uuid.New().String()
	db, err := database.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	repository := repositories.NewPeopleRepository(db)
	err = repository.Insert(people)
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

func SearchByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uuid := params["id"]

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

func TermSearch(w http.ResponseWriter, r *http.Request) {
	term := r.URL.Query().Get("t")

	db, err := database.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	repository := repositories.NewPeopleRepository(db)
	peoples, err := repository.TermSearch(term)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(peoples); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Count(w http.ResponseWriter, r *http.Request) {
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