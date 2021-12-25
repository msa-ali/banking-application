package app

import (
	"log"
	"net/http"
)

func Start() {
	// define routes
	http.HandleFunc("/greet", Greet)
	http.HandleFunc("/customers", GetAllCustomers)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
