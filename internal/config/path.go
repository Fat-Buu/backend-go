package config

import (
	"log"
	"os"
)

func UsersFilePath() string {
	var filePath string = os.Getenv("RESOURCES_USER_JSON")
	log.Println("filePath: ", filePath)
	return filePath
}
