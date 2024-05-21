package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/exoplanets", createExoplanetHandler).Methods("POST")
	r.HandleFunc("/exoplanets", listExoplanetsHandler).Methods("GET")
	r.HandleFunc("/exoplanets/{id}", getExoplanetHandler).Methods("GET")
	r.HandleFunc("/exoplanets/{id}", updateExoplanetHandler).Methods("PUT")
	r.HandleFunc("/exoplanets/{id}", deleteExoplanetHandler).Methods("DELETE")
	r.HandleFunc("/exoplanets/{id}/fuel", fuelEstimationHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
