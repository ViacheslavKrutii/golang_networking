package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type city string

const geocodingAPIkey = "b1b13befe92c375511053642f2310582"
const currentWeatherAPIkey = "75c94e3d992d6c922e7821d4d397ce4b"

type weatherRequest struct {
	City city `json:"city"`
}

type coordinate struct {
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lon"`
}

type geocodingResponse []coordinate

func weather(w http.ResponseWriter, r *http.Request) {
	log.Println("Weather handler")

	body, err := io.ReadAll(r.Body)

	// http body err check
	if err != nil {
		log.Println(err)
		http.Error(w, "internal error1", http.StatusInternalServerError)
		return
	}

	var newWeatherRequest weatherRequest

	err = json.Unmarshal(body, &newWeatherRequest)

	// unmarshal err check
	if err != nil {
		log.Println(err)
		http.Error(w, "internal error2", http.StatusInternalServerError)
		return
	}

	geocodingCityURL := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&appid=%s", newWeatherRequest.City, geocodingAPIkey)

	geocodingCityResponse, err := http.Get(geocodingCityURL)

	if err != nil {
		fmt.Println("City search error:", err)
		return
	}
	defer geocodingCityResponse.Body.Close()

	geocodingCityResponseBody, err := io.ReadAll(geocodingCityResponse.Body)

	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	var newCoordinates geocodingResponse

	err = json.Unmarshal(geocodingCityResponseBody, &newCoordinates)

	if err != nil {
		log.Println(err)
		http.Error(w, "internal error3", http.StatusInternalServerError)
		return
	}

	if len(newCoordinates) == 0 {
		// Обробка, коли координати не знайдено
		http.Error(w, "Coordinates not found", http.StatusNotFound)
		return
	}

	currentWeatherURL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s", newCoordinates[0].Latitude, newCoordinates[0].Longitude, currentWeatherAPIkey)
	curentWeatherResponse, err := http.Get(currentWeatherURL)
	if err != nil {
		fmt.Println("Weather request error:", err)
		return
	}
	defer curentWeatherResponse.Body.Close()

	curentWeatherResponseBody, err := io.ReadAll(curentWeatherResponse.Body)

	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(curentWeatherResponseBody)

}
