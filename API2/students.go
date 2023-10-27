package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type student struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	ClassNum classNum `json:"classNum"`
}

var studentsDB = map[string]student{
	"1": {ID: "1", Name: "John", ClassNum: 5},
	"2": {ID: "2", Name: "Alice", ClassNum: 5},
}

func GetStudentByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["id"]

	for _, student := range studentsDB {
		if student.ID == studentID {
			w.WriteHeader(http.StatusOK)
			responce, _ := json.MarshalIndent(student, "", " ")
			w.Write(responce)
			return
		}
	}

	// Якщо учень не знайдений, повертаємо статус 404 (Not Found)
	w.WriteHeader(http.StatusNotFound)
}
