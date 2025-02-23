package handler

import (
	"fmt"
	"net/http"
)

func EmptyHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "text/html")

	output := "This service does not provide any functionality in the root. Please use <a href=\"" +
		LOCATIONS_PATH + "\">" + LOCATIONS_PATH + "</a> or <a href=\"" + COLLECTION_PATH +
		"\">" + COLLECTION_PATH + "</a>"

	_, err := fmt.Fprint(w, output)

	if err != nil {
		http.Error(w, "Error whne returning output", http.StatusInternalServerError)
	}

}
