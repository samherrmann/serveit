package security

import (
	"os"
	"path/filepath"
)

func fileExists(dir string, filename string) (bool, error) {
	_, err := os.Stat(filepath.Join(dir, filename))
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
