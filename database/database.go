package database

import (
	"database/sql"
	"fmt"
	"root/root/errors"
)

const (
	host     = "localhost"
	port     = 5432
	user     = ""
	password = ""
	dbname   = "postgres"
)

var DB *sql.DB

func InitDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	errors.CheckError(err)
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
