package main

import (
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"
)

var (
	exoplanets = make(map[string]Exoplanet)
	mutex      = &sync.Mutex{}
)

func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func addExoplanet(e Exoplanet) {
	mutex.Lock()
	defer mutex.Unlock()
	exoplanets[e.ID] = e
}

func getAllExoplanets() []Exoplanet {
	mutex.Lock()
	defer mutex.Unlock()
	all := make([]Exoplanet, 0, len(exoplanets))
	for _, exoplanet := range exoplanets {
		all = append(all, exoplanet)
	}
	return all
}

func getExoplanetByID(id string) (Exoplanet, error) {
	mutex.Lock()
	defer mutex.Unlock()
	exoplanet, exists := exoplanets[id]
	if !exists {
		return exoplanet, errors.New("exoplanet not found")
	}
	return exoplanet, nil
}

func updateExoplanet(e Exoplanet) error {
	mutex.Lock()
	defer mutex.Unlock()
	_, exists := exoplanets[e.ID]
	if !exists {
		return errors.New("exoplanet not found")
	}
	exoplanets[e.ID] = e
	return nil
}

func deleteExoplanet(id string) error {
	mutex.Lock()
	defer mutex.Unlock()
	_, exists := exoplanets[id]
	if !exists {
		return errors.New("exoplanet not found")
	}
	delete(exoplanets, id)
	return nil
}

func filterExoplanetsByType(exoplanets []Exoplanet, filterType string) []Exoplanet {
	filtered := make([]Exoplanet, 0)
	for _, exoplanet := range exoplanets {
		if string(exoplanet.Type) == filterType {
			filtered = append(filtered, exoplanet)
		}
	}
	return filtered
}

func sortExoplanets(exoplanets []Exoplanet, sortBy string) {
	switch sortBy {
	case "name":
		sort.SliceStable(exoplanets, func(i, j int) bool {
			return exoplanets[i].Name < exoplanets[j].Name
		})
	case "distance":
		sort.SliceStable(exoplanets, func(i, j int) bool {
			return exoplanets[i].Distance < exoplanets[j].Distance
		})
	case "radius":
		sort.SliceStable(exoplanets, func(i, j int) bool {
			return exoplanets[i].Radius < exoplanets[j].Radius
		})
	case "mass":
		sort.SliceStable(exoplanets, func(i, j int) bool {
			return exoplanets[i].Mass < exoplanets[j].Mass
		})
	}
}
