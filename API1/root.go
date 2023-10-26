package main

import (
	"log"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	log.Println("Served request")
	log.Println(r.Header.Get("User-Agent"))
}
