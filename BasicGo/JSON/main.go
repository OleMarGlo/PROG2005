package main

import (
	"JSON-Example/handler"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("8080")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}

	http.HandleFunc(handler.DEFAULT_PATH, handler.EmptyHandler)
	http.HandleFunc(handler.LOCATIONS_PATH, handler.LocationHandler)
	http.HandleFunc(handler.COLLECTION_PATH, handler.CollectionHandler)

	log.Println("Sarting server on port: " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
