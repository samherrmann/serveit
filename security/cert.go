package security

import (
	"fmt"
	"io/ioutil"
	"net"
	"os/exec"
	"path/filepath"
	"strings"
)

// CertFilename is the server certificate filename.
var CertFilename = "serveit.crt"

// CSRFilename is the server certificate signing request filename.
var CSRFilename = "serveit.csr"

// ExtFilename is the server certificate extensions filename.
var ExtFilename = "serveit.ext"

// EnsureCert creates a server X.509 certificate if it doesn't already exist.
func EnsureCert(dir string, hostnames []string) error {
	exists, err := fileExists(dir, CertFilename)
	if err != nil {
		return err
	}
	if !exists {
		if err := createCSR(dir); err != nil {
			return err
		}
		if err = createExtFile(dir, hostnames); err != nil {
			return err
		}
		return CreateCert(dir)
	}
	return nil
}

// CreateCert creates a server X.509 certificate.
func CreateCert(dir string) error {
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
	cmd.Dir = dir
	_, err := cmd.CombinedOutput()
	return err
}

// createCSR creates a certificate signing request.
func createCSR(dir string) error {
	cmd := exec.Command(
		"openssl", "req",
		"-new",
		"-key", KeyFilename,
		"-subj", "/C=CA/ST=Ontario/L=Ottawa/O=samherrmann/CN=serveit",
		"-out", CSRFilename,
	)
	cmd.Dir = dir
	_, err := cmd.CombinedOutput()
	return err
}

// createExtFile creates a certificate extensions file.
func createExtFile(dir string, hostnames []string) error {
	dns := []string{}
	ips := []string{}

	for _, hostname := range hostnames {
		parsed := net.ParseIP(hostname)
		if parsed != nil {
			l := len(ips)
			ips = append(ips, fmt.Sprintf("IP.%v = %v", l+1, hostname))
		} else {
			l := len(dns)
			dns = append(dns, fmt.Sprintf("DNS.%v = %v", l+1, hostname))
		}
	}

	content := string(
		"authorityKeyIdentifier=keyid,issuer\n" +
			"basicConstraints=CA:FALSE\n" +
			"keyUsage = digitalSignature, nonRepudiation, keyEncipherment, dataEncipherment\n" +
			"subjectAltName = @alt_names\n" +
			"[alt_names]\n",
	)

	content += (strings.Join(dns, "\n") + "\n")
	content += (strings.Join(ips, "\n") + "\n")

	return ioutil.WriteFile(filepath.Join(dir, ExtFilename), []byte(content), 0644)
}
