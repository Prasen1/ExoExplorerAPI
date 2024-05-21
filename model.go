package main

import (
	"errors"
)

type ExoplanetType string

const (
	GasGiant    ExoplanetType = "GasGiant"
	Terrestrial ExoplanetType = "Terrestrial"
)

type Exoplanet struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Distance    float64       `json:"distance"`
	Radius      float64       `json:"radius"`
	Mass        float64       `json:"mass,omitempty"`
	Type        ExoplanetType `json:"type"`
}

func (e *Exoplanet) Validate() error {
	if e.Name == "" || e.Description == "" {
		return errors.New("name and description are required")
	}
	if e.Distance <= 10 || e.Distance >= 1000 {
		return errors.New("distance must be between 10 and 1000 light years")
	}
	if e.Radius <= 0.1 || e.Radius >= 10 {
		return errors.New("radius must be between 0.1 and 10 Earth-radius units")
	}
	if e.Type != GasGiant && e.Type != Terrestrial {
		return errors.New("invalid exoplanet type")
	}
	if e.Type == Terrestrial {
		if e.Mass <= 0.1 || e.Mass >= 10 {
			return errors.New("mass must be between 0.1 and 10 Earth-mass units for terrestrial planets")
		}
	} else {
		e.Mass = 0 // Ensure mass is not set for Gas Giant
	}
	return nil
}
