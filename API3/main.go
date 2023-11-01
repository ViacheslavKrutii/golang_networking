package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/weather", weather).Methods("GET")
	http.ListenAndServe("localhost:8080", r)
}
