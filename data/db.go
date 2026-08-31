package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "postgres"
	PASSWORD = "postgres"
	DBNAME   = "midaas"
)

func NewDB() error {
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASSWORD, DBNAME,
	)

	var dbErr error
	DB, dbErr = sql.Open("postgres", connString)
	if dbErr != nil {
		log.Fatal(dbErr)
		return dbErr
	}

	if err := DB.Ping(); err != nil {
		DB.Close()
		return err
	}

	return nil
}
