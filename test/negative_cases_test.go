package test

import (
	"bytes"
	"encoding/json"
	"exoplanets/handlers"
	"exoplanets/models"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameFieldShouldAllowOnlyAlphabets(t *testing.T) {

	exoplanet := models.ExoPlanetDTO{
		Name:          "A Exoplanet123 hsbasjsksb hjsjhasakskas  jhasashasjash",
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

	if assert.Error(t, handlers.Add(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestDescriptionFieldShouldAllowOnlyAlphabets(t *testing.T) {

	exoplanet := models.ExoPlanetDTO{
		Name:          "A Exoplanet",
		Description:   "A Exoplanet description123",
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

	if assert.Error(t, handlers.Add(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestDistanceFieldShouldAllowOnlyNumericAndHaveMinMaxValue(t *testing.T) {

	exoplanet := models.ExoPlanetDTO{
		Name:          "A Exoplanet",
		Description:   "A Exoplanet description",
		Distance:      10000,
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

	if assert.Error(t, handlers.Add(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestMassFieldShouldAllowOnlyNumericAndHaveMinMaxValue(t *testing.T) {

	exoplanet := models.ExoPlanetDTO{
		Name:          "A Exoplanet",
		Description:   "A Exoplanet description",
		Distance:      1,
		Mass:          11,
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

	if assert.Error(t, handlers.Add(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestRadiusFieldShouldAllowOnlyNumericAndHaveMinMaxValue(t *testing.T) {

	exoplanet := models.ExoPlanetDTO{
		Name:          "A Exoplanet",
		Description:   "A Exoplanet description",
		Distance:      10,
		Mass:          10,
		Radius:        11,
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

	if assert.Error(t, handlers.Add(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestTypeFieldShouldAllowTerrestrialAsValue(t *testing.T) {

	exoplanet := models.ExoPlanetDTO{
		Name:          "A Exoplanet",
		Description:   "A Exoplanet description",
		Distance:      10,
		Mass:          10,
		Radius:        .8,
		ExoplanetType: "Terrestrial",
	}

	payload, err := json.Marshal(exoplanet)
	if err != nil {
		t.Error("inavalid payload")
		return
	}
	body := bytes.NewBuffer(payload)
	c, _, err := GetEchoContext(http.MethodPost, "/v1/api/exoplanets", body)
	if err != nil {
		t.Error(err.Error())
		return
	}
	assert.NoError(t, handlers.Add(c))
}

func TestTypeFieldShouldAllowGasGiantAsValue(t *testing.T) {

	exoplanet := models.ExoPlanetDTO{
		Name:          "A Exoplanet",
		Description:   "A Exoplanet description",
		Distance:      10,
		Mass:          10,
		Radius:        .8,
		ExoplanetType: "GasGiant",
	}

	payload, err := json.Marshal(exoplanet)
	if err != nil {
		t.Error("inavalid payload")
		return
	}
	body := bytes.NewBuffer(payload)
	c, _, err := GetEchoContext(http.MethodPost, "/v1/api/exoplanets", body)
	if err != nil {
		t.Error(err.Error())
		return
	}
	assert.NoError(t, handlers.Add(c))
}

func TestTypeFieldShouldNotAllowOtherValues(t *testing.T) {

	exoplanet := models.ExoPlanetDTO{
		Name:          "A Exoplanet",
		Description:   "A Exoplanet description",
		Distance:      10,
		Mass:          10,
		Radius:        .8,
		ExoplanetType: "Others exoplanent",
	}

	payload, err := json.Marshal(exoplanet)
	if err != nil {
		t.Error("inavalid payload")
		return
	}
	body := bytes.NewBuffer(payload)
	c, _, err := GetEchoContext(http.MethodPost, "/v1/api/exoplanets", body)
	if err != nil {
		t.Error(err.Error())
		return
	}
	assert.Error(t, handlers.Add(c))
}
