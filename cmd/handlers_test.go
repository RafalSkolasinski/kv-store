package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetState(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	app := application{state: &State{}}

	if assert.NoError(t, app.getState(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetValue(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	app := application{state: &State{RequestCount: 42}}

	if assert.NoError(t, app.getState(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var state State
		err := json.Unmarshal(rec.Body.Bytes(), &state)
		require.NoError(t, err)
		assert.Equal(t, 42, int(state.RequestCount))
	}
}

func TestPutValue(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPut, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	app := application{state: &State{}}

	if assert.NoError(t, app.putValueHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "not implemented")
	}
}

func TestDeleteValue(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	app := application{state: &State{}}

	if assert.NoError(t, app.deleteValueHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "not implemented")
	}
}
