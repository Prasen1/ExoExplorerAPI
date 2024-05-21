package main

import (
	"errors"
	"strconv"
)

func estimateFuel(e Exoplanet, crewCapacityStr string) (float64, error) {
	crewCapacity, err := strconv.Atoi(crewCapacityStr)
	if err != nil || crewCapacity <= 0 {
		return 0, errors.New("invalid crew capacity")
	}
	var g float64
	if e.Type == GasGiant {
		g = 0.5 / (e.Radius * e.Radius)
	} else if e.Type == Terrestrial {
		g = e.Mass / (e.Radius * e.Radius)
	} else {
		return 0, errors.New("invalid exoplanet type")
	}
	fuel := e.Distance / (g * g) * float64(crewCapacity)
	return fuel, nil
}
