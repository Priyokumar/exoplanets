package fuel

import (
	"errors"
	"exoplanets/gravity"
	"exoplanets/repository"
	"math"
	"net/http"
)

func EstimateFuel(id int64, capacity int) (float64, int, error) {
	planet, code, err := repository.GetByID(id)
	if err != nil {
		return 0.0, code, err
	}
	var g float64
	if planet.ExoplanetType == "GasGiant" {
		g, err = gravity.CalculateGravity(gravity.GasExoplanetGravity{Planet: *planet})
		if err != nil {
			return 0.0, http.StatusInternalServerError, err
		}
	} else if planet.ExoplanetType == "Terrestrial" {
		g, err = gravity.CalculateGravity(gravity.TerrestrialExoplanetGravity{Planet: *planet})
		if err != nil {
			return 0.0, http.StatusInternalServerError, err
		}
	} else {
		return 0.0, http.StatusNotFound, errors.New("unknown exoplanet")
	}

	fuel := float64(planet.Distance) / ((math.Pow(g, 2)) * float64(capacity))
	return fuel, http.StatusOK, nil

}
