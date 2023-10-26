package main

import (
	"errors"
	"log"
	"net/http"
)

func authorize(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Println("Chacking authorization")

		cookie, err := req.Cookie("session")

		if err != nil {
			switch {
			case errors.Is(err, http.ErrNoCookie):
				http.Error(w, "You are not logged in", http.StatusBadRequest)
			default:
				http.Error(w, "Server error", http.StatusInternalServerError)
			}
			return
		}
		if !authenticated[login(cookie.Value)] {
			http.Error(w, "your session is expired", http.StatusUnauthorized)
		}

		fn.ServeHTTP(w, req)

	}

}
