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
type role string
type classNum uint8
type name string

type registerForm struct {
	Login    login    `json:"login"`
	Password password `json:"password"`
	Role     role     `json:"role"`
	ClassNum classNum `json:"classNum"`
	Name     name     `json:"name"`
}

var registered map[login]password = make(map[login]password)
var permisionList map[login]bool = make(map[login]bool)

func register(w http.ResponseWriter, r *http.Request) {
	log.Println("Register handler")
	body, err := io.ReadAll(r.Body)

	// http body err check
	if err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	var newRegisterForm registerForm

	err = json.Unmarshal(body, &newRegisterForm)

	// unmarshal err check
	if err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	// empty data check
	if newRegisterForm.Login == "" || newRegisterForm.Password == "" {
		http.Error(w, "login and password must be present", http.StatusBadRequest)
	}
	// uniqe check
	if _, exists := registered[newRegisterForm.Login]; exists {
		http.Error(w, "login already taken", http.StatusBadRequest)
		return
	}

	//

	h := sha1.New()
	h.Write([]byte(newRegisterForm.Password))
	registered[newRegisterForm.Login] = password(hex.EncodeToString(h.Sum(nil)))

	if newRegisterForm.Role != "teacher" {
		permisionList[newRegisterForm.Login] = false
	} else {
		permisionList[newRegisterForm.Login] = true
	}

	log.Printf("%+v\n", registered)

	w.Write([]byte("You are registered"))
}
