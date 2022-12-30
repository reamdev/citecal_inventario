package util

import (
	"citecal_inventario/storage/logs"
	"log"
	"os"
	"strconv"
)

func GetEnvStringValue(envParam string) string {
	value, err := os.LookupEnv(envParam)

	if !err {
		errorMessage := "No se encontro valor para " + envParam
		logs.WriteInLogFile(errorMessage, logFileName)
		log.Fatalln(errorMessage)
	}

	return value
}

func GetEnvIntValue(envParam string) int64 {
	value := GetEnvStringValue(envParam)

	intValue, convErr := strconv.ParseInt(value, 0, 64)

	if convErr != nil {
		errorMessage := "El valor de " + envParam + " no se puede convertir a int: " + convErr.Error()
		logs.WriteInLogFile(errorMessage, logFileName)
		log.Fatalln(errorMessage)
	}

	return intValue
}
