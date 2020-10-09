package security

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

// ReadCertFile returns the x.509v3 certificate from the file named by the
// provided filename.
func ReadCertFile(filename string) (*x509.Certificate, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("cannot read %v: %w", filename, err)
	}

	block, _ := pem.Decode(bytes)

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("cannot parse %v: %w", filename, err)
	}
	return cert, nil
}
