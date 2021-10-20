package common

import (
	"fmt"
	"os"
)

func CreateDirectoryIfNotExists(relPath string) {
	path, _ := os.Getwd()

	if _, err := os.Stat(fmt.Sprintf("%s/%s", path, relPath)); os.IsNotExist(err) {
		_ = os.Mkdir(relPath, os.ModePerm)
	}
}
