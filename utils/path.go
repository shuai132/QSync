package utils

import (
	"os"
	"path/filepath"
)

func GetFullPath(filePath string) string {
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return filepath.Join(path, filePath)
}