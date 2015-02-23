package main

import (
	"log"
	"net/http"

	"../routes"
)

// Starts the webserver
func main() {
	router := routes.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
