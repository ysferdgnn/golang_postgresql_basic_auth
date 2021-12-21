package services

import (
	"fmt"
	"postgres/database/handlers"
	"postgres/database/models"
)

func GetAllUsers() []models.User {

	user_list := make([]models.User, 1)

	database, err := handlers.SetupDB()

	if err == nil {
		rows, errorQ := database.Query("SELECT * FROM kullanici")

		if errorQ == nil {
			for rows.Next() {
				var id int
				var name string

				rows.Scan(&id, &name)

				user_tmp := models.User{Id: id, Name: name}

				user_list = append(user_list, user_tmp)
			}
		} else {
			fmt.Println(errorQ.Error())
		}

	} else {
		fmt.Println(err.Error())
	}

	return user_list

}
