package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"postgres/database/services"
	"strings"
)

func HandleUser(w http.ResponseWriter, r *http.Request) {

	var username string
	var password string
	for name, val := range r.Header {

		// fmt.Println(fmt.Sprintf("Headers. name: %s, value: %s", name, val))

		if name == "Username" {
			username = strings.Join(val, "")

		}
		if name == "Password" {
			password = strings.Join(val, "")
		}
	}
	fmt.Println(fmt.Sprintf("Request arrived. username: %s, Password: %s", username, password))

	if username == "Saruman" && password == "theWhite" {

		users := services.GetAllUsers()
		usersBytes := new(bytes.Buffer)
		for _, item := range users {
			json.NewEncoder(usersBytes).Encode(item)
		}

		w.Write(usersBytes.Bytes())
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("It is Forbidden Bro!"))
	}
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var username string
	var password string
	for name, val := range r.Header {

		if name == "Username" {
			username = strings.Join(val, "")

		}
		if name == "Password" {
			password = strings.Join(val, "")
		}
	}
	fmt.Println(fmt.Sprintf("Request arrived. username: %s, Password: %s", username, password))

	if username == "Saruman" && password == "theWhite" {

		services.AddUser()
		w.Write([]byte("User Added"))
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("It is Forbidden Bro!"))
	}
}
