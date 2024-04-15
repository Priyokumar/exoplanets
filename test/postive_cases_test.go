package test

import (
	"bytes"
	"encoding/json"
	"exoplanets/handlers"
	"exoplanets/models"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var id int64

func TestAAdd(t *testing.T) {

	exoplanet := models.ExoPlanetDTO{
		Name:          "A Exoplanet",
		Description:   "A Exoplanet description",
		Distance:      123,
		Mass:          0.2,
		Radius:        0.5,
		ExoplanetType: "Terrestrial",
	}

	payload, err := json.Marshal(exoplanet)
	if err != nil {
		t.Error("inavalid payload")
		return
	}
	body := bytes.NewBuffer(payload)
	c, rec, err := GetEchoContext(http.MethodPost, "/v1/api/exoplanets", body)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if assert.NoError(t, handlers.Add(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.NotNil(t, rec.Body, rec.Body)
		var exoplanet models.ExoPlanet
		if assert.NoError(t, json.NewDecoder(io.Reader(rec.Body)).Decode(&exoplanet)) {
			assert.NotZero(t, exoplanet.ID)
			id = exoplanet.ID
		}

	}

}

func TestBGet(t *testing.T) {
	c, rec, err := GetEchoContext(http.MethodGet, "/v1/api/exoplanets", nil)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if assert.NoError(t, handlers.Get(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotNil(t, rec.Body, rec.Body)
		var exoplanets = []models.ExoPlanet{}
		if assert.NoError(t, json.NewDecoder(io.Reader(rec.Body)).Decode(&exoplanets)) {
			assert.NotEmpty(t, exoplanets)
		}

	}

}

func TestCGetByID(t *testing.T) {
	c, rec, err := GetEchoContext(http.MethodGet, "/v1/api/exoplanets", nil)
	if err != nil {
		t.Error(err.Error())
		return
	}
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprint(id))
	if assert.NoError(t, handlers.GetByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotNil(t, rec.Body, rec.Body)
		var exoplanet models.ExoPlanet
		if assert.NoError(t, json.NewDecoder(io.Reader(rec.Body)).Decode(&exoplanet)) {
			assert.NotEmpty(t, exoplanet)
		}
	}

}

func TestDUpdate(t *testing.T) {

	exoplanet := models.ExoPlanetDTO{
		Name:          "A Exoplanet B",
		Description:   "A Exoplanet description",
		Distance:      123,
		Mass:          0.2,
		Radius:        0.5,
		ExoplanetType: "Terrestrial",
	}

	payload, err := json.Marshal(exoplanet)
	if err != nil {
		t.Error("inavalid payload")
		return
	}
	body := bytes.NewBuffer(payload)
	c, rec, err := GetEchoContext(http.MethodPost, "/v1/api/exoplanets", body)
	if err != nil {
		t.Error(err.Error())
		return
	}
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprint(id))
	if assert.NoError(t, handlers.Update(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotNil(t, rec.Body, rec.Body)
		var exoplanet models.ExoPlanet
		if assert.NoError(t, json.NewDecoder(io.Reader(rec.Body)).Decode(&exoplanet)) {
			assert.NotZero(t, exoplanet.ID)
			id = exoplanet.ID
		}
		assert.Equal(t, "A Exoplanet B", exoplanet.Name)
	}
}

func TestEGEstimatedFuel(t *testing.T) {
	c, rec, err := GetEchoContext(http.MethodGet, "/v1/api/exoplanets", nil)
	if err != nil {
		t.Error(err.Error())
		return
	}
	c.SetParamNames("id", "capacity")
	c.SetParamValues(fmt.Sprint(id), fmt.Sprint(456))
	if assert.NoError(t, handlers.EstimateFuel(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.NotNil(t, rec.Body, rec.Body)
		var fuel map[string]float64
		if assert.NoError(t, json.NewDecoder(io.Reader(rec.Body)).Decode(&fuel)) {
			assert.NotEmpty(t, fuel)
		}
		assert.Greater(t, fuel["estimatedFuel"], 0.0)
	}
}
