package gravity

import (
	"exoplanets/models"
	"math"
)

type GasExoplanetGravity struct {
	Planet models.ExoPlanet
}

func (g GasExoplanetGravity) Calculate() (float64, error) {
	return (0.5 / math.Pow(g.Planet.Radius, 2)), nil
}
