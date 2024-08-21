package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type application struct {
	state *State
	store *Store
}

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	app := application{state: &State{}, store: NewStore()}
	app.registerRoutes(e)

	e.Logger.Fatal(e.Start("localhost:8080"))
}
