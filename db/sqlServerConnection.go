package db

import (
	_ "github.com/denisenkom/go-mssqldb"

	"citecal_inventario/config"
	"citecal_inventario/storage/logs"
	"context"
	"database/sql"
	"fmt"
	"log"
)

var Connection *sql.DB

func InitSQLServerConnection() {
	var err error
	sqlConfig := config.ConfigParams.SQLServerParams

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", sqlConfig.Server, sqlConfig.User, sqlConfig.Password, sqlConfig.Port, sqlConfig.Database)

	Connection, err = sql.Open("sqlserver", connString)

	if err != nil {
		errorMessage := "Error al crear la conexion: " + err.Error()

		logs.WriteInLogFile(errorMessage, logFileName)
		log.Fatalln(errorMessage)
	}

	log.Println("Conectado a SQL Server en el puerto", sqlConfig.Port)

	defer Connection.Close()

	getSQLServerVersion()
}

func getSQLServerVersion() {
	// Use background context
	ctx := context.Background()

	// Ping database to see if it's still alive.
	// Important for handling network issues and long queries.
	err := Connection.PingContext(ctx)
	if err != nil {
		errorMessage := "Error pinging database: " + err.Error()
		logs.WriteInLogFile(errorMessage, logFileName)
		log.Fatalln(errorMessage)
	}

	var result string

	// Run query and scan for result
	err = Connection.QueryRowContext(ctx, "SELECT @@version").Scan(&result)
	if err != nil {
		errorMessage := "Scan failed:" + err.Error()
		logs.WriteInLogFile(errorMessage, logFileName)
		log.Fatalln(errorMessage)
	}

	log.Printf("\n%s\n", result)
}
