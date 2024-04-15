package repository

import (
	"exoplanets/models"
	"net/http"
	"slices"
)

var id int64
var data = []models.ExoPlanet{}

func Get(filterData map[string]string, sortBy string) ([]models.ExoPlanet, int, error) {
	if len(data) < 1 {
		return nil, http.StatusNotFound, nil
	} else {
		exoplanets := slices.Clone(data)
		exoplanets = filter(exoplanets, getCriterias(filterData)...)
		exoplanets = sortData(exoplanets, sortBy)
		return exoplanets, http.StatusOK, nil
	}
}

func Add(d *models.ExoPlanet) (*models.ExoPlanet, int, error) {
	id++
	d.ID = id
	if slices.ContainsFunc(data, func(ep models.ExoPlanet) bool {
		return ep.Name == d.Name
	}) {
		return nil, http.StatusConflict, nil
	}
	data = append(data, *d)

	return d, http.StatusCreated, nil
}

func GetByID(id int64) (*models.ExoPlanet, int, error) {
	i := slices.IndexFunc(data, func(ep models.ExoPlanet) bool {
		return ep.ID == id
	})
	if i > -1 {
		return &data[i], http.StatusOK, nil
	} else {
		return nil, http.StatusNotFound, nil
	}
}

func Update(id int64, d *models.ExoPlanet) (*models.ExoPlanet, int, error) {
	i := slices.IndexFunc(data, func(ep models.ExoPlanet) bool {
		return ep.ID == id
	})
	if i > -1 {
		d.ID = data[i].ID
		data[i] = *d
		return d, http.StatusOK, nil
	} else {
		return nil, http.StatusNotFound, nil
	}
}

func Delete(id int64) (int, error) {
	i := slices.IndexFunc(data, func(ep models.ExoPlanet) bool {
		return ep.ID == id
	})
	if i > -1 {
		data = slices.Delete(data, i, i+1)
		return http.StatusOK, nil
	} else {
		return http.StatusNotFound, nil
	}
}
