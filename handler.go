package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func createExoplanetHandler(w http.ResponseWriter, r *http.Request) {
	var exoplanet Exoplanet
	if err := json.NewDecoder(r.Body).Decode(&exoplanet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := exoplanet.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	exoplanet.ID = generateID()
	addExoplanet(exoplanet)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(exoplanet)
	log.Printf("Added planet with Id: %s", exoplanet.ID)
}

func listExoplanetsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	sortBy := query.Get("sortBy")
	filterType := query.Get("type")

	exoplanets := getAllExoplanets()

	if filterType != "" {
		exoplanets = filterExoplanetsByType(exoplanets, filterType)
	}

	if sortBy != "" {
		sortExoplanets(exoplanets, sortBy)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exoplanets)
	log.Printf("Retrieved %d planet records", len(exoplanets))
}

func getExoplanetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	exoplanet, err := getExoplanetByID(id)
	if err != nil {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exoplanet)
	log.Printf("Retrieved planet with Id: %s", exoplanet.ID)
}

func updateExoplanetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var exoplanet Exoplanet
	if err := json.NewDecoder(r.Body).Decode(&exoplanet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	exoplanet.ID = id
	if err := exoplanet.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := updateExoplanet(exoplanet); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(exoplanet)
	log.Printf("Updated planet with Id: %s", exoplanet.ID)
}

func deleteExoplanetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if err := deleteExoplanet(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	log.Printf("Deleted planet with Id: %s", id)
}

func fuelEstimationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	crewCapacityStr := r.URL.Query().Get("crewCapacity")
	if crewCapacityStr == "" {
		http.Error(w, "crewCapacity is required", http.StatusBadRequest)
		return
	}
	crewCapacity, err := strconv.Atoi(crewCapacityStr)
	if err != nil || crewCapacity <= 0 {
		http.Error(w, "invalid crewCapacity", http.StatusBadRequest)
		return
	}
	exoplanet, err := getExoplanetByID(id)
	if err != nil {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}
	fuel, err := estimateFuel(exoplanet, crewCapacityStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{"fuel": fuel})
	log.Printf("Returned fuel estimate for planet Id: %s", exoplanet.ID)
}
