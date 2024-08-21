package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type envelope map[string]any

type GetRequest struct {
	Key string `param:"key"`
}

type PutRequest struct {
	GetRequest
	Value string `json:"value"`
}

type DeleteRequest struct {
}

func (app *application) getState(c echo.Context) error {
	c.Logger().Debugf("state: %+v", *app.state)
	return c.JSON(http.StatusOK, app.state)
}

func (app *application) getValueHandler(c echo.Context) error {
	var input struct {
		InputKey
	}

	if err := c.Bind(&input); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "bad request")
	}
	c.Logger().Debugf("processing get value request for key `%s`", input.Key)

	value, err := app.store.Get(input.Key)
	if err != nil {
		if errors.Is(err, ErrNoKey) {
			return c.JSON(http.StatusNotFound, envelope{"error": err})
		}
		return c.JSON(http.StatusInternalServerError, envelope{"error": "server error"})
	}

	return c.JSON(http.StatusOK, envelope{"value": value})
}

func (app *application) putValueHandler(c echo.Context) error {
	var input struct {
		InputKey
		InputValue
	}

	if err := c.Bind(&input); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "bad request")
	}
	c.Logger().Debugf("processing put value request for (`%s`, `%s`)", input.Key, input.Value)

	err := app.store.Put(input.Key, input.Value)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, envelope{"error": err})
	}

	return c.JSON(http.StatusOK, envelope{"value": "added to store"})
}

func (app *application) deleteValueHandler(c echo.Context) error {
	var input struct {
		InputKey
	}

	if err := c.Bind(&input); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "bad request")
	}
	c.Logger().Debugf("processing delete value request for key `%s`", input.Key)

	err := app.store.Delete(input.Key)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, envelope{"error": err})
	}

	return c.JSON(http.StatusOK, envelope{"value": "deleted from store"})
}
