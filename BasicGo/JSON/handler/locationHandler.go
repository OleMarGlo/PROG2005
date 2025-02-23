package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func LocationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handlePostRequest(w, r)
	case http.MethodGet:
		handleGetRequest(w, r)
	default:
		http.Error(w, "REST method: '"+r.Method+"' not supported. Current only '"+http.MethodGet+"' and '"+http.MethodPost+"' Are supported",
			http.StatusNotImplemented)
	}
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Recived POST request")

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	Location := Location{}

	err := decoder.Decode(&Location)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("Received following request: ", Location)

}

func handleGetRequest(w http.ResponseWriter, r *http.Request) {

	location := Location{
		Name:     "Kebab",
		Postcode: "2900",
		Geolocation: Coordinates{
			Latitude:   61.100,
			Longtitude: 10.700,
		},
	}

	w.Header().Set("content-type", "application/json")

	encoder := json.NewEncoder(w)

	err := encoder.Encode(location)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
