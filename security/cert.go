package security

import (
	"io/ioutil"
	"os"
	"os/exec"
)

// CertFilename is the application certificate filename.
var CertFilename = "serveit.crt"

// CSRFilename is the application certificate signing request filename.
var CSRFilename = "serveit.csr"

// ExtFilename is the application certificate extensions filename.
var ExtFilename = "serveit.ext"

// EnsureCert creates an application X.509 certificate if it doesn't already
// exist.
func EnsureCert() error {
	_, err := os.Stat(CertFilename)
	if os.IsNotExist(err) {
		if err := createCSR(); err != nil {
			return err
		}
		if err = createExtFile(); err != nil {
			return err
		}
		return CreateCert()
	}
	return err
}

// CreateCert creates an application X.509 certificate.
func CreateCert() error {
	cmd := exec.Command(
		"openssl", "x509",
		"-req",
		"-in", CSRFilename,
		"-CA", RootCACertFilename,
		"-CAkey", RootCAKeyFilename,
		"-passin", "pass:serveit",
		"-CAcreateserial",
		"-out", CertFilename,
		"-days", "3650",
		"-sha256",
		"-extfile", ExtFilename,
	)
	_, err := cmd.CombinedOutput()
	return err
}

// createCSR creates a certificate signing request.
func createCSR() error {
	cmd := exec.Command(
		"openssl", "req",
		"-new",
		"-key", KeyFilename,
		"-subj", "/C=CA/ST=Ontario/L=Ottawa/O=samherrmann/CN=serveit",
		"-out", CSRFilename,
	)
	_, err := cmd.CombinedOutput()
	return err
}

// createExtFile creates a certificate extensions file.
func createExtFile() error {
	content := []byte(
		"authorityKeyIdentifier=keyid,issuer\n" +
			"basicConstraints=CA:FALSE\n" +
			"keyUsage = digitalSignature, nonRepudiation, keyEncipherment, dataEncipherment\n" +
			"subjectAltName = @alt_names\n" +
			"[alt_names]\n" +
			"DNS.1 = localhost\n",
	)
	return ioutil.WriteFile(ExtFilename, content, 0644)
}
