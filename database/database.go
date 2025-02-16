package database

import (
	"database/sql"
	"fmt"
	// "os"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	stringConnection := fmt.Sprintf("root:liedsonfsa@/rinha?charset=utf8&parseTime=True&loc=Local")
	// fmt.Println(stringConnection)
	db, err := sql.Open("mysql", stringConnection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}