package security

import (
	"os/exec"
)

var fileRootname = "serveit_root_ca"

// RootCACertFilename is the certificate authority (CA) X.509 certificate filename.
var RootCACertFilename = fileRootname + ".crt"

// RootCACertSerialFilename is the certificate authority (CA) serial number file.
var RootCACertSerialFilename = fileRootname + ".srl"

// EnsureRootCACert creates a certificate authority (CA) X.509 certificate if it
// doesn't already exist.
func EnsureRootCACert(dir string) error {
	exists, err := fileExists(dir, RootCACertFilename)
	if err != nil {
		return err
	}
	if !exists {
		return CreateRootCACert(dir)
	}
	return nil
}

// CreateRootCACert creates a certificate authority (CA) X.509 certificate.
func CreateRootCACert(dir string) error {
	cmd := exec.Command(
		"openssl", "req",
		"-x509",
		"-new",
		"-nodes",
		"-key", RootCAKeyFilename,
		"-sha256",
		"-days", "3650",
		"-passin", "pass:serveit",
		"-subj", `/C=CA/ST=Ontario/L=Ottawa/O=samherrmann/CN=Serveit Root Certificate Authority`,
		"-out", RootCACertFilename,
	)
	cmd.Dir = dir
	_, err := cmd.CombinedOutput()
	return err
}
