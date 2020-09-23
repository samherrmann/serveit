package execpath

import (
	"os"
	"path/filepath"
)

// Dir returns the absolute path of the directory that contains the executable
// that started the current process. The returned path does not include a
// trailing slash.
func Dir() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(exe), nil
}
