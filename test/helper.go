package test

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

func GetEchoContext(method string, URL string, body io.Reader) (echo.Context, *httptest.ResponseRecorder, error) {
	e := echo.New()
	req, err := http.NewRequest(method, URL, body)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c, rec, nil
}
