package repository

import (
	"exoplanets/models"
	"sort"
	"strconv"
)

type Exoplanets []models.ExoPlanet

func filter(data []models.ExoPlanet, criterias ...func(models.ExoPlanet) bool) []models.ExoPlanet {

	if len(criterias) < 1 {
		return data
	}
	var result Exoplanets
	for _, value := range data {
		okCriteria := true
		for _, cr := range criterias {
			if !cr(value) {
				okCriteria = false
			}
		}
		if okCriteria {
			result = append(result, value)
		}
	}
	return result
}

func getCriterias(filterData map[string]string) []func(models.ExoPlanet) bool {
	criterias := make([]func(models.ExoPlanet) bool, 0)
	if _, ok := filterData["name"]; ok {
		criterias = append(criterias, func(ep models.ExoPlanet) bool {
			return ep.Name == filterData["name"]
		})
	}
	if _, ok := filterData["mass"]; ok {
		criterias = append(criterias, func(ep models.ExoPlanet) bool {
			mass, _ := strconv.ParseFloat(filterData["mass"], 64)
			return ep.Mass == mass
		})
	}
	return criterias
}

func sortData(data []models.ExoPlanet, sortBy string) []models.ExoPlanet {
	if sortBy == "mass" {
		massSortingFunc := func(i, j int) bool {
			return data[i].Mass > data[j].Mass
		}
		sort.Slice(data, massSortingFunc)
	}
	if sortBy == "radius" {
		radiusSortingFunc := func(i, j int) bool {
			return data[i].Radius > data[j].Radius
		}
		sort.Slice(data, radiusSortingFunc)
	}
	return data
}
