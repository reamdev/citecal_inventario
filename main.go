package main

import (
	"citecal_inventario/config"
	"citecal_inventario/db"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Cargar env File
	config.LoadEnvConfig()
	db.InitSQLServerConnection()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	// Iniciar Aplicacion en puerto definido
	port := fmt.Sprintf(":%d", config.ConfigParams.Port)
	e.Logger.Fatal(e.Start(port))
}
