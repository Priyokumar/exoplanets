package gravity

import (
	"exoplanets/models"
	"math"
)

type TerrestrialExoplanetGravity struct {
	Planet models.ExoPlanet
}

func (g TerrestrialExoplanetGravity) Calculate() (float64, error) {
	return (g.Planet.Mass / math.Pow(g.Planet.Radius, 2)), nil
}
