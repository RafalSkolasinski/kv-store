package main

import (
	"sync/atomic"

	"github.com/labstack/echo/v4"
)

// RequestCounterMiddleware is a middleware that counts incoming requests
func (app *application) RequestCounterMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		atomic.AddUint64(&app.state.RequestCount, 1)
		return next(c)
	}
}
