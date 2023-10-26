package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var authenticated map[login]bool = make(map[login]bool)

func authenticate(w http.ResponseWriter, r *http.Request) {
	log.Println("authenticate handler")
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

	originPassword, loginExist := registered[newAttempt.Login]
	if !loginExist {
		http.Error(w, "login doesn`t exist", http.StatusBadRequest)
		return
	}

	h := sha1.New()
	h.Write([]byte(newAttempt.Password))
	guessPassword := password(hex.EncodeToString(h.Sum(nil)))

	if originPassword != guessPassword {
		http.Error(w, "wrong pass", http.StatusBadRequest)
		return
	}

	authenticated[newAttempt.Login] = true

	cookie := http.Cookie{
		Name:     "session",
		Value:    string(newAttempt.Login),
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(w, &cookie)

	w.Write([]byte("you are logged in"))
}
