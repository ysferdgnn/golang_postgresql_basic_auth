package services

import (
	"fmt"
	"postgres/database/handlers"
	"postgres/database/models"
	"time"
)

func GetAllUsers() []models.User {

	user_list := make([]models.User, 1)

	database, err := handlers.SetupDB()

	if err == nil {
		database.AutoMigrate(&models.User{})

		database.Model(&models.User{}).Take(&user_list)

	} else {
		fmt.Print(err.Error())
		panic(err.Error())
	}

	return user_list

}
func AddUser() {
	database, err := handlers.SetupDB()

	if err == nil {
		database.AutoMigrate(&models.User{})

		user := models.User{Name: "Yusuf", CreatedAt: time.Now(), UpdatedAt: time.Now()}
		database.Model(&models.User{}).Save(&user)

	} else {
		fmt.Print(err.Error())
		panic(err.Error())
	}

}
