package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type Pessoa struct {
	UUID 			string		`json:"id"`
	Name 			string		`json:"nome"`
	Apelido 		string		`json:"apelido"`
	Nascimento 		string		`json:"nascimento"`
	Stack 			[]string	`json:"stack"`
}

func main() {
	http.HandleFunc("/pessoas", postPessoas)

	http.ListenAndServe(":3000", nil)
}

func postPessoas(w http.ResponseWriter, r* http.Request) {
	var pessoa Pessoa
	if err := json.NewDecoder(r.Body).Decode(&pessoa); err != nil {
		fmt.Println(err)
		return
	}

	pessoa.UUID = uuid.New().String()

	fmt.Println(pessoa)
}
