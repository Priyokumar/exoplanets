package models

import (
	"errors"
	"fmt"
	"regexp"
)

type ExoPlanet struct {
	ID            int64   `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Distance      int64   `json:"distance"`
	Mass          float64 `json:"mass"`
	Radius        float64 `json:"radius"`
	ExoplanetType string  `json:"exoplanetType"`
	Metadata
}

type Metadata struct {
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type ExoPlanetDTO struct {
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Distance      int64   `json:"distance"`
	Mass          float64 `json:"mass"`
	Radius        float64 `json:"radius"`
	ExoplanetType string  `json:"exoplanetType"`
}

func (e ExoPlanet) Validate() error {

	// Name should allow only alphabetic characters and min 1and max 50 chanracters
	if !regexp.MustCompile(`^[a-zA-Z\s]{1,50}$`).MatchString(e.Name) {
		return errors.New("name field allows only alphabetic characters, min length 1 and max length 50")
	}

	if !regexp.MustCompile(`^[a-zA-Z\s]{1,500}$`).MatchString(e.Description) {
		return errors.New("description field allows only alphabetic characters, min length 1 and max length 50")
	}

	if !regexp.MustCompile(`^\d`).MatchString(fmt.Sprint(e.Distance)) || (e.Distance < 10 || e.Distance > 1000) {
		return errors.New("distance field allows only numeric characters, min length 10 and max length 1000")
	}

	if !regexp.MustCompile(`^\d`).MatchString(fmt.Sprint(e.Mass)) || (float32(e.Mass) < 0.1 || float32(e.Mass) > 10) {
		return errors.New("mass field allows only numeric characters, min length 0.1 and max length 10")
	}

	if !regexp.MustCompile(`^\d`).MatchString(fmt.Sprint(e.Radius)) || (float32(e.Radius) < 0.1 || float32(e.Radius) > 10) {
		return errors.New("radius field allows only numeric characters, min length 10 and max length 1000")
	}

	if !regexp.MustCompile(`\b(GasGiant|Terrestrial)\b`).MatchString(e.ExoplanetType) {
		return errors.New("unsopported exoplanet type. exoplanetType field allows either GasGiant or Terrestrial")
	}

	return nil
}
