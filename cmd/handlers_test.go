package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetState(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	e := echo.New()
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

func TestGetValue(t *testing.T) {
	values := map[string]string{"test-key": "test-value"}

	req := httptest.NewRequest(http.MethodGet, "/values/test-key", nil)
	rec := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, rec)

	c.SetParamNames("key")
	c.SetParamValues("test-key")

	app := application{store: &Store{values: values}}
	if assert.NoError(t, app.getValueHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Equal(t, "test-value", response["data"])
	}
}

func TestPutValue(t *testing.T) {
	payload := `{"value": "test-value"}`

	req := httptest.NewRequest(http.MethodPut, "/values/test-key", strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, rec)

	c.SetParamNames("key")
	c.SetParamValues("test-key")

	app := application{store: NewStore()}
	if assert.NoError(t, app.putValueHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		require.NoError(t, err)

		assert.Equal(t, "added to store", response["data"])
		assert.Equal(t, app.store.values["test-key"], "test-value")
	}
}

func TestDeleteValue(t *testing.T) {
	values := map[string]string{"test-key": "test-value"}

	req := httptest.NewRequest(http.MethodDelete, "/values/test-key", nil)
	rec := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, rec)

	c.SetParamNames("key")
	c.SetParamValues("test-key")

	app := application{store: &Store{values: values}}
	if assert.NoError(t, app.deleteValueHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response map[string]string
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		require.NoError(t, err)
		assert.Equal(t, "deleted from store", response["data"])
		assert.NotContains(t, app.store.values, "test-key")
	}
}
