package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
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
var idCounterTeacher = 2
var idCounterStudent = 3

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

	//registerForm validation
	shouldReturn := validation(newRegisterForm, w)
	if shouldReturn {
		return
	}

	//permision set
	setPermision(newRegisterForm)

	switch newRegisterForm.Role {
	case "teacher":
		var newTeacher teacher
		err = json.Unmarshal(body, &newTeacher)
		if err != nil {
			log.Println(err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
		newTeacher.ID = strconv.Itoa(idCounterTeacher)
		idCounterTeacher++
		addTeacherDB(newTeacher)
		addTeacherToClass(newTeacher)

	case "student":
		var newStudent student
		err = json.Unmarshal(body, &newStudent)
		if err != nil {
			log.Println(err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
		newStudent.ID = strconv.Itoa(idCounterStudent)
		idCounterStudent++
		addStudentDB(newStudent)
		addStudentToClass(newStudent)
	}

	h := sha1.New()
	h.Write([]byte(newRegisterForm.Password))
	registered[newRegisterForm.Login] = password(hex.EncodeToString(h.Sum(nil)))

	log.Printf("%+v\n", registered)

	w.Write([]byte("You are registered"))
}

func setPermision(newRegisterForm registerForm) {
	if newRegisterForm.Role == "teacher" {
		permisionList[newRegisterForm.Login] = true
	} else if newRegisterForm.Role == "student" {
		permisionList[newRegisterForm.Login] = false
	}
}

func validation(newRegisterForm registerForm, w http.ResponseWriter) bool {
	// empty data check
	if newRegisterForm.Login == "" || newRegisterForm.Password == "" {
		http.Error(w, "login and password must be present", http.StatusBadRequest)
		return true
	}
	// uniqe check
	if _, exists := registered[newRegisterForm.Login]; exists {
		http.Error(w, "login already taken", http.StatusBadRequest)
		return true
	}

	// role check
	if newRegisterForm.Role != "teacher" && newRegisterForm.Role != "student" {
		http.Error(w, "Wrong role", http.StatusBadRequest)
		return true
	}

	//class exist check
	if _, exists := classes[newRegisterForm.ClassNum]; !exists {
		http.Error(w, "Class not found", http.StatusBadRequest)
		return true
	}

	return false
}
