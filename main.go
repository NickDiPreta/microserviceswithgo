package main

import (
	"log"
	"net/http"
)

func main() {

	// Anything other than 'goodbye'
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Hello World")
	})

	// Function called when goodbye path is called
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})

	http.ListenAndServe(":9090", nil)
}
