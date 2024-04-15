package routes

import (
	"exoplanets/handlers"

	"github.com/labstack/echo/v4"
)

func Set(e *echo.Echo) {
	e.Static("/", "web")
	prefix := "/v1/api/exoplanets"
	e.GET(prefix, handlers.Get)
	e.POST(prefix, handlers.Add)
	e.GET(prefix+"/:id", handlers.GetByID)
	e.PUT(prefix+"/:id", handlers.Update)
	e.DELETE(prefix+"/:id", handlers.Delete)
	e.GET(prefix+"/:id/estimatedfuel", handlers.EstimateFuel)
}
