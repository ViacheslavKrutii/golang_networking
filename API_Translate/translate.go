package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/bregydoc/gtranslate"
)

type userText struct {
	Text string `json:"text"`
	From string `json:"from"`
	To   string `json:"to"`
}

func translate(w http.ResponseWriter, r *http.Request) {
	log.Println("Translate handler")
	body, err := io.ReadAll(r.Body)

	// http body err check
	if err != nil {
		log.Println(err)
		http.Error(w, "internal error1", http.StatusInternalServerError)
		return
	}

	var newUserText userText

	err = json.Unmarshal(body, &newUserText)

	// unmarshal err check
	if err != nil {
		log.Println(err)
		http.Error(w, "internal error2", http.StatusInternalServerError)
		return
	}

	translated, err := gtranslate.TranslateWithParams(
		newUserText.Text,
		gtranslate.TranslationParams{
			From: newUserText.From,
			To:   newUserText.To,
		},
	)
	if err != nil {
		panic(err)
	}

	w.Write([]byte(translated))
}
