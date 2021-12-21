package services

import (
	"fmt"
	"postgres/database/handlers"
	"postgres/database/models"
)

func GetAllUsers() []models.Kullanici {

	kullanici_listesi := make([]models.Kullanici, 1)

	database, err := handlers.SetupDB()

	if err == nil {
		rows, errorQ := database.Query("SELECT * FROM kullanici")

		if errorQ == nil {
			for rows.Next() {
				var id int
				var ad string

				rows.Scan(&id, &ad)

				kullanici_tmp := models.Kullanici{Id: id, Ad: ad}

				kullanici_listesi = append(kullanici_listesi, kullanici_tmp)
			}
		} else {
			fmt.Println(errorQ.Error())
		}

	} else {
		fmt.Println(err.Error())
	}

	return kullanici_listesi

}
