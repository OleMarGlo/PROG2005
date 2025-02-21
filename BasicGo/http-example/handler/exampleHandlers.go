package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func DiagHandlerComplex(w http.ResponseWriter, r *http.Request) {

	output := "Complex handler" + "\n"

	output += "Request:" + "\n"
	output += "URL path: " + r.URL.Path + "\n"
	output += "Path Value: " + r.PathValue("param1") + "\n"
	output += "Path Value: " + r.PathValue("param2") + "\n"
	output += "Method: " + r.Method + "\n"

	_, err := fmt.Fprintf(w, "%v", output)
	if err != nil {
		log.Print("An error occured: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
