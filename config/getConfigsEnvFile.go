package config

import (
	"citecal_inventario/storage/logs"
	"citecal_inventario/util"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

// Iniciar estructura de configuracion
var ConfigParams ServerConfig

// Funcion para cargar configuraciones de entorno (.env) en estructura ServerConfig
func LoadEnvConfig() {
	// Cargar env file
	err := godotenv.Load()

	if err != nil {
		errorMessage := "Error al cargar env: " + err.Error()
		log.Printf(errorMessage)
		logs.WriteInLogFile(errorMessage, logFileName)
	}

	// Guardar parametros de SQL Server en structure
	sqlSConfig := SQLServerConfig{
		Server:   util.GetEnvStringValue("SQLSERVER_SERVER"),
		Port:     uint16(util.GetEnvIntValue("SQLSERVER_PORT")),
		User:     util.GetEnvStringValue("SQLSERVER_USER"),
		Password: util.GetEnvStringValue("SQLSERVER_PASSWORD"),
		Database: util.GetEnvStringValue("SQLSERVER_DATABASE"),
	}

	ConfigParams = ServerConfig{
		Port:            uint16(util.GetEnvIntValue("PORT")),
		SQLServerParams: sqlSConfig,
	}
}
