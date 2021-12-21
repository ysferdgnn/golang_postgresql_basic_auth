package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"postgres/database/services"
	"strings"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/users", HandleMain)

	http.ListenAndServe("localhost:8081", mux)

}

func HandleMain(w http.ResponseWriter, r *http.Request) {

	var username string
	var password string
	for name, val := range r.Header {

		fmt.Println(fmt.Sprintf("Header geldi. ad: %s, değer: %s", name, val))

		if name == "Username" {
			username = strings.Join(val, "")
		}
		if name == "Password" {
			password = strings.Join(val, "")
		}
	}
	fmt.Println(fmt.Sprintf("İstek geldi. Kullanici: %s, Şifre: %s", username, password))

	if username == "yusuf" && password == "123" {

		kullanicilar := services.GetAllUsers()
		kullanicilarBytes := new(bytes.Buffer)
		for _, item := range kullanicilar {
			json.NewEncoder(kullanicilarBytes).Encode(item)
		}

		w.Write(kullanicilarBytes.Bytes())
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Forbidden kardeşim"))
	}
}
