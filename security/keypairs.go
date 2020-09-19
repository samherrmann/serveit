package security

import (
	"fmt"
)

// EnsureKeyPairs creates an RSA key and a X.509 certificate for both the
// certificate authority (CA) and the server if they don't already exist.
func EnsureKeyPairs(hostnames []string) error {
	err := EnsureRootCAKey()
	if err != nil {
		return fmt.Errorf("Error creating %v: %w", RootCAKeyFilename, err)
	}
	err = EnsureRootCACert()
	if err != nil {
		return fmt.Errorf("Error creating %v: %w", RootCACertFilename, err)
	}
	err = EnsureKey()
	if err != nil {
		return fmt.Errorf("Error creating %v: %w", KeyFilename, err)
	}
	err = EnsureCert(hostnames)
	if err != nil {
		return fmt.Errorf("Error creating %v: %w", CertFilename, err)
	}
	return nil
}
