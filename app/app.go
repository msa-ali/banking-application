package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// define routes
	router.HandleFunc("/greet", Greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", GetCustomer).Methods(http.MethodGet) // Only matches when customer id is numeric value otherwise 404 error

	router.HandleFunc("/customers", CreateCustomer).Methods((http.MethodPost))

	router.HandleFunc("/api/time", GetCurrentTime)
	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
