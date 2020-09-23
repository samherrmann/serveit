package security

import (
	"os/exec"
)

// KeyFilename is the server RSA key filename.
var KeyFilename = "serveit.key"

// EnsureKey creates an RSA key if it doesn't already exist.
func EnsureKey(dir string) error {
	exists, err := fileExists(dir, KeyFilename)
	if err != nil {
		return err
	}
	if !exists {
		return CreateKey(dir)
	}
	return nil
}

// CreateKey creates an RSA key.
func CreateKey(dir string) error {
	cmd := exec.Command(
		"openssl", "genrsa",
		"-out", KeyFilename,
		"2048",
	)
	cmd.Dir = dir
	_, err := cmd.CombinedOutput()
	return err
}
