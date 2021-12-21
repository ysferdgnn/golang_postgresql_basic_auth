package handlers

import (
	"database/sql"
	"fmt"
	"postgres/database/constants"

	_ "github.com/lib/pq"
)

func SetupDB() (*sql.DB, error) {

	connection_string := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", constants.DB_USER, constants.DB_PASSWORD, constants.DB_NAME)

	database, err := sql.Open("postgres", connection_string)

	return database, err

}
