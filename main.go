package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	details "github.com/ksalim-k/go-microservices/details"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Checking the health")
	response := map[string]string{
		"status":    "UP",
		"timestamp": time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving the homepage")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "We are in the homepage!!!!")
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching the details")
	hostname, err := details.GetHostName()
	if err != nil {
		log.Fatal(err)
	}
	ip, _ := details.GetIP()
	response := map[string]string{
		"hostname": hostname,
		"IP":       ip.String(),
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/details", getHandler)
	log.Println("Server started...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
