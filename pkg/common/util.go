package common

import (
	"fmt"
	"os"
)

func createDirectoryIfNotExists(rel_path string) {
	path, _ := os.Getwd()

	if _, err := os.Stat(fmt.Sprintf("%s/%s", path, rel_path)); os.IsNotExist(err) {
		_ = os.Mkdir(rel_path, os.ModePerm)
	}
}
