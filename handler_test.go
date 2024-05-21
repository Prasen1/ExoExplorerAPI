package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func setupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/exoplanets", createExoplanetHandler).Methods("POST")
	r.HandleFunc("/exoplanets", listExoplanetsHandler).Methods("GET")
	r.HandleFunc("/exoplanets/{id}", getExoplanetHandler).Methods("GET")
	r.HandleFunc("/exoplanets/{id}", updateExoplanetHandler).Methods("PUT")
	r.HandleFunc("/exoplanets/{id}", deleteExoplanetHandler).Methods("DELETE")
	r.HandleFunc("/exoplanets/{id}/fuel", fuelEstimationHandler).Methods("GET")
	return r
}

func TestCreateExoplanetHandler(t *testing.T) {
	router := setupRouter()

	exoplanet := Exoplanet{
		Name:        "TestPlanet",
		Description: "A test planet",
		Distance:    100,
		Radius:      1,
		Type:        Terrestrial,
		Mass:        1,
	}
	body, _ := json.Marshal(exoplanet)

	req, err := http.NewRequest("POST", "/exoplanets", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	var createdExoplanet Exoplanet
	if err := json.NewDecoder(rr.Body).Decode(&createdExoplanet); err != nil {
		t.Errorf("handler returned invalid JSON: %v", err)
	}

	if createdExoplanet.ID == "" {
		t.Errorf("handler did not assign an ID")
	}
}

func TestListExoplanetsHandler(t *testing.T) {
	router := setupRouter()

	req, err := http.NewRequest("GET", "/exoplanets", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var exoplanets []Exoplanet
	if err := json.NewDecoder(rr.Body).Decode(&exoplanets); err != nil {
		t.Errorf("handler returned invalid JSON: %v", err)
	}
}

func TestGetExoplanetHandler(t *testing.T) {
	router := setupRouter()

	exoplanet := Exoplanet{
		ID:          "test-id",
		Name:        "TestPlanet",
		Description: "A test planet",
		Distance:    100,
		Radius:      1,
		Type:        Terrestrial,
		Mass:        1,
	}
	addExoplanet(exoplanet)

	req, err := http.NewRequest("GET", "/exoplanets/test-id", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var returnedExoplanet Exoplanet
	if err := json.NewDecoder(rr.Body).Decode(&returnedExoplanet); err != nil {
		t.Errorf("handler returned invalid JSON: %v", err)
	}

	if returnedExoplanet.ID != exoplanet.ID {
		t.Errorf("handler returned wrong exoplanet: got %v want %v",
			returnedExoplanet.ID, exoplanet.ID)
	}
}

func TestUpdateExoplanetHandler(t *testing.T) {
	router := setupRouter()

	exoplanet := Exoplanet{
		ID:          "test-id",
		Name:        "TestPlanet",
		Description: "A test planet",
		Distance:    100,
		Radius:      1,
		Type:        Terrestrial,
		Mass:        1,
	}
	addExoplanet(exoplanet)

	updatedExoplanet := Exoplanet{
		Name:        "UpdatedPlanet",
		Description: "An updated planet",
		Distance:    200,
		Radius:      2,
		Type:        Terrestrial,
		Mass:        2,
	}
	body, _ := json.Marshal(updatedExoplanet)

	req, err := http.NewRequest("PUT", "/exoplanets/test-id", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var returnedExoplanet Exoplanet
	if err := json.NewDecoder(rr.Body).Decode(&returnedExoplanet); err != nil {
		t.Errorf("handler returned invalid JSON: %v", err)
	}

	if returnedExoplanet.Name != updatedExoplanet.Name {
		t.Errorf("handler returned wrong exoplanet: got %v want %v",
			returnedExoplanet.Name, updatedExoplanet.Name)
	}
}

func TestDeleteExoplanetHandler(t *testing.T) {
	router := setupRouter()

	exoplanet := Exoplanet{
		ID:          "test-id",
		Name:        "TestPlanet",
		Description: "A test planet",
		Distance:    100,
		Radius:      1,
		Type:        Terrestrial,
		Mass:        1,
	}
	addExoplanet(exoplanet)

	req, err := http.NewRequest("DELETE", "/exoplanets/test-id", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}

	if _, err := getExoplanetByID("test-id"); err == nil {
		t.Errorf("exoplanet was not deleted")
	}
}

func TestFuelEstimationHandler(t *testing.T) {
	router := setupRouter()

	exoplanet := Exoplanet{
		ID:          "test-id",
		Name:        "TestPlanet",
		Description: "A test planet",
		Distance:    100,
		Radius:      1,
		Type:        Terrestrial,
		Mass:        1,
	}
	addExoplanet(exoplanet)

	req, err := http.NewRequest("GET", "/exoplanets/test-id/fuel?crewCapacity=10", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response map[string]float64
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("handler returned invalid JSON: %v", err)
	}

	if _, ok := response["fuel"]; !ok {
		t.Errorf("handler did not return fuel estimation")
	}
}
