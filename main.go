package main

import (
	"net/http"
	api "postgres/api/handlers"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/users", api.HandleUser)
	mux.HandleFunc("/users/add", api.AddUser)

	http.ListenAndServe("localhost:8081", mux)

}
