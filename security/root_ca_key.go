package security

import (
	"os/exec"
)

// RootCAKeyFilename is the certificate authority (CA) RSA key filename.
var RootCAKeyFilename = "serveit_root_ca.key"

// EnsureRootCAKey creates a certificate authority (CA) RSA key if it doesn't
// already exist.
func EnsureRootCAKey(dir string) error {
	exists, err := fileExists(dir, RootCAKeyFilename)
	if err != nil {
		return err
	}
	if !exists {
		return CreateRootCAKey(dir)
	}
	return nil
}

// CreateRootCAKey creates a certificate authority (CA) RSA key.
func CreateRootCAKey(dir string) error {
	cmd := exec.Command(
		"openssl", "genrsa",
		"-aes256",
		"-passout", "pass:serveit",
		"-out", RootCAKeyFilename,
		"2048",
	)
	cmd.Dir = dir
	_, err := cmd.CombinedOutput()
	return err
}
