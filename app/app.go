package app

import (
	"log"
	"net/http"
)

func Start() {

	mux := http.NewServeMux()

	// define routes
	mux.HandleFunc("/greet", Greet)
	mux.HandleFunc("/customers", GetAllCustomers)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
