package security

import (
	"os"
	"os/exec"
)

var fileRootname = "serveit_root_ca"

// RootCACertFilename is the certificate authority (CA) X.509 certificate filename.
var RootCACertFilename = fileRootname + ".crt"

// RootCACertSerialFilename is the certificate authority (CA) serial number file.
var RootCACertSerialFilename = fileRootname + ".srl"

// EnsureRootCACert creates a certificate authority (CA) X.509 certificate if it
// doesn't already exist.
func EnsureRootCACert() error {
	_, err := os.Stat(RootCACertFilename)
	if os.IsNotExist(err) {
		return CreateRootCACert()
	}
	return err
}

// CreateRootCACert creates a certificate authority (CA) X.509 certificate.
func CreateRootCACert() error {
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
	_, err := cmd.CombinedOutput()
	return err
}
