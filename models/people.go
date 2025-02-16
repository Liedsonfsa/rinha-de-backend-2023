package models

type People struct {
	UUID 			string		`json:"id"`
	Name 			string		`json:"nome"`
	Apelido 		string		`json:"apelido"`
	Nascimento 		string		`json:"nascimento"`
	Stack 			[]string	`json:"stack,omitempty"`
}