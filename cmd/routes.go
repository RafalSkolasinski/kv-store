package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (app *application) registerRoutes(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.Use(app.RequestCounterMiddleware)

	e.GET("/state", app.getState)

	e.GET("/values/:key", app.getValueHandler)
	e.PUT("/values/:key", app.putValueHandler)
	e.DELETE("/values/:key", app.deleteValueHandler)
}
