package logs

import (
	"fmt"
	"log"
	"os"
)

// Function for write errors in logs files
func WriteInLogFile(message string, logType string) {
	logFileRoute := fmt.Sprintf("./storage/logs/%s.log", logType)

	logFile, err := os.OpenFile(logFileRoute, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal("Error al abrir archivo: ", err.Error())
	}

	defer logFile.Close()

	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println(message)
}
