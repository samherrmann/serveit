package security

import (
	"os"
	"os/exec"
)

// RootCAKeyFilename is the certificate authority (CA) RSA key filename.
var RootCAKeyFilename = "serveit_root_ca.key"

// EnsureRootCAKey creates a certificate authority (CA) RSA key if it doesn't
// already exist.
func EnsureRootCAKey() error {
	_, err := os.Stat(RootCAKeyFilename)
	if os.IsNotExist(err) {
		return CreateRootCAKey()
	}
	return err
}

// CreateRootCAKey creates a certificate authority (CA) RSA key.
func CreateRootCAKey() error {
	cmd := exec.Command(
		"openssl", "genrsa",
		"-aes256",
		"-passout", "pass:serveit",
		"-out", RootCAKeyFilename,
		"2048",
	)
	_, err := cmd.CombinedOutput()
	return err
}
