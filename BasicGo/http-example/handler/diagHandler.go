package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func DiagHandler(w http.ResponseWriter, r *http.Request) {
	output := "Request:" + "\n"
	output += "URL path: " + r.URL.Path + "\n"
	output += "Path Value: " + r.PathValue("param1") + "\n"
	output += "Method: " + r.Method + "\n"

	output += "Headers: " + "\n"
	for k, v := range r.Header {
		for _, vv := range v {
			output += k + ": " + vv + "\n"
		}
	}

	//body very crude
	content, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Error during body extraction: " + err.Error())
		http.Error(w, "Error during decoding", http.StatusInternalServerError)
		return
	}
	output += "\n" + "Content:" + "\n"
	output += string(content)

	_, err = fmt.Fprintf(w, "%v", output)
	if err != nil {
		log.Print("An error occured: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
