package util

import (
	"citecal_inventario/storage/logs"

	"golang.org/x/crypto/bcrypt"
)

func EncryptValue(value string) (string, error) {
	var hashString string
	valueByteArr := []byte(value)

	hash, err := bcrypt.GenerateFromPassword(valueByteArr, bcrypt.DefaultCost)

	if err != nil {
		errorMessage := "Error al generar Hash: " + err.Error()
		logs.WriteInLogFile(errorMessage, logFileName)
	}

	hashString = string(hash)

	return hashString, err
}
