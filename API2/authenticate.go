package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type authenticateForm struct {
	Login    login    `json:"login"`
	Password password `json:"password"`
}

var authenticated map[login]bool = make(map[login]bool)

func authenticate(w http.ResponseWriter, r *http.Request) {
	log.Println("authenticate handler")
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	var newAuthenticateForm authenticateForm

	err = json.Unmarshal(body, &newAuthenticateForm)

	if err != nil {
		log.Println(err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	originPassword, loginExist := registered[newAuthenticateForm.Login]
	if !loginExist {
		http.Error(w, "login doesn`t exist", http.StatusBadRequest)
		return
	}

	h := sha1.New()
	h.Write([]byte(newAuthenticateForm.Password))
	guessPassword := password(hex.EncodeToString(h.Sum(nil)))

	if originPassword != guessPassword {
		http.Error(w, "wrong pass", http.StatusBadRequest)
		return
	}

	authenticated[newAuthenticateForm.Login] = true

	cookieAuth := http.Cookie{
		Name:     "session",
		Value:    string(newAuthenticateForm.Login),
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteNoneMode,
	}

	// hash := sha1.New()
	// hash.Write([]byte("true"))
	// valueCookiePermission := hex.EncodeToString(hash.Sum(nil))

	// cookiePermission := http.Cookie{
	// 	Name:     "permission",
	// 	Value:    valueCookiePermission,
	// 	Path:     "/",
	// 	MaxAge:   3600,
	// 	HttpOnly: true,
	// 	Secure:   false,
	// 	SameSite: http.SameSiteNoneMode,
	// }

	// if permisionList[newAuthenticateForm.Login] {
	// 	http.SetCookie(w, &cookiePermission)
	// } else {
	// 	http.Error(w, "you have no permission", http.StatusBadRequest)
	// }

	http.SetCookie(w, &cookieAuth)

	w.Write([]byte("you are logged in"))
}
