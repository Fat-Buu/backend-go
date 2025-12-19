package config

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func UsersFilePath() string {
	wd, _ := os.Getwd()
	var l []string = strings.Split(wd, "\\")
	log.Println("Working dir: ", wd)
	var filePath string
	if l[len(l)-1] == "test" {
		filePath = filepath.Join(wd, "..", "resources", "users.json")
	} else {
		filePath = filepath.Join(wd, "resources", "users.json")

	}
	log.Println("filePath: ", filePath)

	return filePath
}
