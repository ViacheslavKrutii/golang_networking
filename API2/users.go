package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var usersDB = []map[string]any{{"name": "Greg"}, {"name": "John"}}

func users(w http.ResponseWriter, r *http.Request) {
	log.Println("users handler")

	w.WriteHeader(http.StatusOK)
	responce, _ := json.MarshalIndent(usersDB, "", " ")
	w.Write(responce)

}
