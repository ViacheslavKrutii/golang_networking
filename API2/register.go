package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type login string
type password string

type attempt struct {
	Login    login    `json:"login"`
	Password password `json:"password"`
}

var registered map[login]password = make(map[login]password)

func register(w http.ResponseWriter, r *http.Request) {
	log.Println("Register handler")
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	var newAttempt attempt

	err = json.Unmarshal(body, &newAttempt)

	if err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	if newAttempt.Login == "" || newAttempt.Password == "" {
		http.Error(w, "login and password must be present", http.StatusBadRequest)
	}

	if _, exists := registered[newAttempt.Login]; exists {
		http.Error(w, "login already taken", http.StatusBadRequest)
		return
	}

	h := sha1.New()
	h.Write([]byte(newAttempt.Password))
	registered[newAttempt.Login] = password(hex.EncodeToString(h.Sum(nil)))

	log.Printf("%+v\n", registered)

	w.Write([]byte("You are registered"))
}
