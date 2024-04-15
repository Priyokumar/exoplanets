package handlers

import (
	"errors"
	"net/url"
	"regexp"
	"strconv"
)

type Error struct {
	Code    int    `json:"code"`
	Summary string `json:"summary"`
}

func validateAndGetFilterQueryParams(qParams url.Values) (map[string]string, error) {
	filterData := map[string]string{}
	if qParams.Has("name") {
		if !regexp.MustCompile(`^[a-zA-Z\s]{1,50}$`).MatchString(qParams.Get("name")) {
			return nil, errors.New("name field allows only alphabetic characters, min length 1 and max length 50")
		}
		filterData["name"] = qParams.Get("name")
	}
	if qParams.Has("mass") {
		if !regexp.MustCompile(`^\d`).MatchString(qParams.Get("mass")) {
			return nil, errors.New("distance field allows only numeric characters")
		}
		mass, _ := strconv.ParseFloat(qParams.Get("mass"), 64)
		if mass < 0.1 || mass > 10 {
			return nil, errors.New("min length 0.1 and max length 10")
		} else {
			filterData["mass"] = qParams.Get("mass")
		}
	}
	return filterData, nil
}

func validateAndGetSortQueryParam(qParams url.Values) (string, error) {
	if qParams.Has("sortBy") {
		if !regexp.MustCompile(`\b(radius|mass)\b`).MatchString(qParams.Get("sortBy")) {
			return "", errors.New("sortBy field should be either radius or mass ")
		}
		return qParams.Get("sortBy"), nil
	}

	return "", nil
}
