package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *application) getState(c echo.Context) error {
	c.Logger().Debugf("state: %+v", *app.state)
	return c.JSON(http.StatusOK, app.state)
}

func (app *application) getValueHandler(c echo.Context) error {
	c.Logger().Debug("processing get value request...")
	return c.JSON(http.StatusOK, "get not implemented...")
}

func (app *application) putValueHandler(c echo.Context) error {
	c.Logger().Debug("processing put value request...")
	return c.JSON(http.StatusOK, "put not implemented...")
}

func (app *application) deleteValueHandler(c echo.Context) error {
	c.Logger().Debug("processing delete value request...")
	return c.JSON(http.StatusOK, "delete not implemented...")
}
