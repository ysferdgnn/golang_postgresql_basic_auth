package handlers

import (
	"fmt"
	"postgres/database/constants"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

func SetupDB() (*gorm.DB, error) {

	// connection_string := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", constants.DB_USER, constants.DB_PASSWORD, constants.DB_NAME)
	connection_string := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable ",
		constants.DB_HOST, constants.DB_USER, constants.DB_PASSWORD, constants.DB_NAME, constants.DB_PORT)
	database, err := gorm.Open(postgres.Open(connection_string), &gorm.Config{})

	return database, err

}
