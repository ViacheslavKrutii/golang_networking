package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var subscriptionsDB = []map[string]any{{"login": []string{"sub1", "sub2"}}}

func subscriptions(w http.ResponseWriter, r *http.Request) {
	log.Println("subscriptions handler")

	w.WriteHeader(http.StatusOK)
	responce, _ := json.MarshalIndent(subscriptionsDB, "", " ")
	w.Write(responce)

}
