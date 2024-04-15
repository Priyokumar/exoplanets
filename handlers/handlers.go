package handlers

import (
	"exoplanets/fuel"
	"exoplanets/models"
	"exoplanets/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {

	filterData, err := validateAndGetFilterQueryParams(c.QueryParams())
	if err != nil {
		log.Println(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, Error{Code: http.StatusBadRequest, Summary: err.Error()})
	}
	sortBy, err := validateAndGetSortQueryParam(c.QueryParams())
	if err != nil {
		log.Println(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, Error{Code: http.StatusBadRequest, Summary: err.Error()})
	}
	data, code, err := repository.Get(filterData, sortBy)
	if err != nil {
		log.Println(err.Error())
		echo.NewHTTPError(code, Error{Code: code, Summary: err.Error()})
	}
	return c.JSON(code, data)
}

func Add(c echo.Context) error {
	var planet models.ExoPlanet
	err := c.Bind(&planet)
	if err != nil {
		log.Println(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, Error{Code: http.StatusBadRequest, Summary: err.Error()})
	}
	err = planet.Validate()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, Error{Code: http.StatusBadRequest, Summary: err.Error()})
		return err
	}
	data, code, err := repository.Add(&planet)
	if err != nil {
		log.Println(err.Error())

		return echo.NewHTTPError(code, Error{Code: code, Summary: err.Error()})
	}
	return c.JSON(code, data)
}

func GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, Error{Code: http.StatusBadRequest, Summary: "invalid id. expected integer value"})
	}
	data, code, err := repository.GetByID(int64(id))
	if err != nil {
		log.Println(err.Error())
		return echo.NewHTTPError(code, Error{Code: code, Summary: err.Error()})
	}
	return c.JSON(code, data)
}

func Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, Error{Code: http.StatusBadRequest, Summary: "invalid id. expected integer value"})
	}
	var planet models.ExoPlanet
	err = c.Bind(&planet)
	if err != nil {
		log.Println(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, Error{Code: http.StatusBadRequest, Summary: err.Error()})
	}
	err = planet.Validate()
	if err != nil {
		log.Println(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, Error{Code: http.StatusBadRequest, Summary: err.Error()})
	}
	data, code, err := repository.Update(int64(id), &planet)
	if err != nil {
		log.Println(err.Error())
		return echo.NewHTTPError(code, Error{Code: code, Summary: err.Error()})
	}
	return c.JSON(code, data)
}

func Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, Error{Code: http.StatusBadRequest, Summary: "invalid id. expected integer value"})
	}
	code, err := repository.Delete(int64(id))
	if err != nil {
		log.Println(err.Error())
		return echo.NewHTTPError(code, Error{Code: code, Summary: err.Error()})
	}
	return c.JSON(code, nil)
}

func EstimateFuel(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, Error{Code: http.StatusBadRequest, Summary: "invalid id. expected integer value"})
	}
	capacity := 2
	if c.QueryParam("capacity") != "" {
		capacity, err = strconv.Atoi(c.QueryParam("capacity"))
		if err != nil {
			log.Println(err.Error())
			return echo.NewHTTPError(http.StatusBadRequest, Error{Code: http.StatusBadRequest, Summary: "invalid capacity. expected integer value"})
		}
	}
	data, code, err := fuel.EstimateFuel(int64(id), capacity)
	if err != nil {
		log.Println(err.Error())
		return echo.NewHTTPError(code, Error{Code: code, Summary: err.Error()})
	}
	return c.JSON(code, map[string]float64{"estimatedFuel": data})
}
