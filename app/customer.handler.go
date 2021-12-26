package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Altamashattari/banking-application/service"
	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!!")
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := ch.service.GetAllCustomers()
	if err != nil {
		writeResponse(w, err.Code, err.AsMesssage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
	// if r.Header.Get("Content-Type") == "application/xml" {
	// 	w.Header().Add("Content-Type", "application/xml")
	// 	xml.NewEncoder(w).Encode(customers)
	// } else {
	// 	w.Header().Add("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(customers)
	// }
}

func (ch *CustomerHandlers) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	customer, err := ch.service.GetCustomer(customerId)
	if err != nil {
		writeResponse(w, err.Code, err.AsMesssage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}

}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post Request Received")
}

func GetCurrentTime(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string, 0)
	tz := r.URL.Query().Get("tz")
	timezones := strings.Split(tz, ",")

	if len(timezones) <= 1 {
		loc, err := time.LoadLocation(tz)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(fmt.Sprintf("invalid timezone %s", tz)))
		} else {
			response["current_time"] = time.Now().In(loc).String()
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}
	} else {
		for _, tzdb := range timezones {
			loc, err := time.LoadLocation(tzdb)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(fmt.Sprintf("invalid timezone %s", tzdb)))
			}
			now := time.Now().In(loc)
			response[tzdb] = now.String()
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
