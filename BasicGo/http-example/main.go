package main

import (
	handlers "http-example/handler"
	"log"
	"net/http"
	"os"
)

func main() {

	PORT := "8080"
	if os.Getenv("PORT") != "" {
		PORT = os.Getenv("PORT")
	}

	router := http.NewServeMux()
	//writing POST before the URL it will only work for the POST method, you should use http.Method... for this
	router.HandleFunc("/diag/{param1}/{param2}", handlers.DiagHandlerComplex) //it can only hadle one extra

	router.HandleFunc("/diag/{param1...}", handlers.DiagHandler) //this can handle all however it is only one param
	//(made for patarns)
	//router.HandleFunc("/diag/", handlers.DiagHandler)					//This can handle all aswell

	log.Printf("%s", "Starting server on port "+PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
