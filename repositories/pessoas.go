package repositories

import (
	"database/sql"
	"encoding/json"
	"rinha-de-backend-2023/models"
)

type PeopleRepository struct {
	db *sql.DB
}

func NewPeopleRepository(db *sql.DB) *PeopleRepository {
	return &PeopleRepository{db}
}

func (peopleRepo PeopleRepository) Insert(people models.People) error {
	stm, err := peopleRepo.db.Prepare("INSERT INTO pessoas (id, apelido, nome, nascimento, stack) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stm.Close()

	stack, err := json.Marshal(people.Stack)
	if err != nil {
		return err
	}
	_, err = stm.Exec(people.UUID, people.Apelido, people.Name, people.Nascimento, stack)
	if err != nil {
		return err
	}

	return nil
}

func (peopleRepo PeopleRepository) Search() ([]models.People, error) {
	rows, err := peopleRepo.db.Query("SELECT * FROM pessoas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var peoples []models.People

	for rows.Next() {
		var people models.People
		var stackJSON sql.NullString

		if err = rows.Scan(&people.UUID, &people.Apelido, &people.Name, &people.Nascimento, &stackJSON); err != nil {
			return nil, err
		}

		if stackJSON.Valid {
			if err = json.Unmarshal([]byte(stackJSON.String), &people.Stack); err != nil {
				return nil, err
			}
		}

		peoples = append(peoples, people)
	}

	return peoples, nil
}