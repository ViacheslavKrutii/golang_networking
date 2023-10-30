package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type task struct {
	Tasktext string `json:"tasktext"`
}

var tasksDB = map[string][]task{
	"26.10.2023": {{Tasktext: "wash the dishes"}, {Tasktext: "walk the dog"}},
	"27.10.2023": {{Tasktext: "buy apartment"}},
}

func tasks(w http.ResponseWriter, r *http.Request) {
	log.Println("tasks handler")

	// body must be dd.mm.yyyy
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	responce, _ := json.MarshalIndent(tasksDB[string(body)], "", " ")
	w.Write(responce)

}
