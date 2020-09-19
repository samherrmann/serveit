package security

import (
	"os"
	"os/exec"
)

// KeyFilename is the server RSA key filename.
var KeyFilename = "serveit.key"

// EnsureKey creates an RSA key if it doesn't already exist.
func EnsureKey() error {
	_, err := os.Stat(KeyFilename)
	if os.IsNotExist(err) {
		return CreateKey()
	}
	return err
}

// CreateKey creates an RSA key.
func CreateKey() error {
	cmd := exec.Command(
		"openssl", "genrsa",
		"-out", KeyFilename,
		"2048",
	)
	_, err := cmd.CombinedOutput()
	return err
}
