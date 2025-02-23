package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func CollectionHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		handleMapPostRequest(w, r)
	case http.MethodGet:
		handleMapGetRequest(w, r)
	default:
		http.Error(w, "REST Method: '"+r.Method+"' not supported. Currently '"+http.MethodGet+
			"' and '"+http.MethodPost+"' supported", http.StatusNotImplemented)
	}
}

func handleMapPostRequest(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	mp := map[string]interface{}{}

	err := decoder.Decode(&mp)
	if err != nil {
		http.Error(w, "Error during decoding", http.StatusBadRequest)
		return
	}
	fmt.Println("Recieved request: ")
	fmt.Println(mp)

	output, err := json.MarshalIndent(mp, "", " ")
	if err != nil {
		http.Error(w, "Error during printing", http.StatusInternalServerError)
	}
	fmt.Println("Pretty printing: ")
	fmt.Println(string(output))

	http.Error(w, "OK", http.StatusOK)
}

func handleMapGetRequest(w http.ResponseWriter, r *http.Request) {

	collection := map[string]interface{}{
		"first":  "firstValue",
		"second": 2,
		"third":  3.1,
	}

	w.Header().Set("content-type", "application/json")

	encoder := json.NewEncoder(w)

	err := encoder.Encode(collection)
	if err != nil {
		http.Error(w, "Error during encoding", http.StatusInternalServerError)
		return
	}

	http.Error(w, "", http.StatusOK)
}
