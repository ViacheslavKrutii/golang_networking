package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", root)
	r.HandleFunc("/users", authorize(users))
	r.HandleFunc("/subscriptions", authorize(subscriptions))

	r.HandleFunc("/register", register)
	r.HandleFunc("/login", authenticate)

	http.ListenAndServe("localhost:8080", r)
}
