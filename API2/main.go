package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", root)
	r.HandleFunc("/users", authorize(users))

	r.HandleFunc("/register", register)
	r.HandleFunc("/login", authenticate)

	r.HandleFunc("/student/{id}", permission(GetStudentByID)).Methods("GET")

	http.ListenAndServe("localhost:8080", r)
}
