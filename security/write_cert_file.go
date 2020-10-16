package security

import (
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net"
	"os"
	"time"
)

// WriteCertFile creates a X.509v3 certificate for the first level in the given
// chain and writes it to the file named by the provided filename. Note that the
// file named by filename must not exist or an error is returned. A file-exist
// error may be checked with errors.Is(err, os.ErrExist). The file is created
// with the given perm mode.
func WriteCertFile(chain *ChainOfTrust, perm os.FileMode) error {
	// Attempt to create and open file. Error if file already exists.
	fileFlag := os.O_WRONLY | os.O_CREATE | os.O_EXCL
	file, err := os.OpenFile(chain.Filename, fileFlag, perm)
	if err != nil {
		return fmt.Errorf("failed to open %v: %w", chain.Filename, err)
	}
	// Create certificate serial number.
	serialNumber, err := CreateSerialNumber()
	if err != nil {
		return fmt.Errorf("failed to generate serial number for %v: %w", chain.Filename, err)
	}

	// Split the hosts list into DNS names and IP addresses.
	dnsNames := []string{}
	ipAddresses := []net.IP{}
	for _, host := range chain.Hosts {
		if ip := net.ParseIP(host); ip != nil {
			ipAddresses = append(ipAddresses, ip)
		} else {
			dnsNames = append(dnsNames, host)
		}
	}

	// Define the certificate template.
	certTemplate := &x509.Certificate{
		SerialNumber:          serialNumber,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(0, 0, chain.Days),
		Subject:               *chain.Subject,
		DNSNames:              dnsNames,
		IPAddresses:           ipAddresses,
		BasicConstraintsValid: !chain.leaf,
		IsCA:                  !chain.leaf,
		ExtKeyUsage:           extKeyUsage(chain.leaf),
		KeyUsage:              keyUsage(chain.leaf),
	}
	// Read the RSA private key file.
	key, err := ReadKeyFile(chain.KeyFilename, chain.KeyPassword)
	if err != nil {
		return fmt.Errorf("failed to read RSA private key file %v: %w", chain.KeyFilename, err)
	}
	// Use the key and certificate of the given chain level also as the parent key
	// and certificate in case the given chain level doesn't have a parent.
	parentKey := key
	parentCertTemplate := certTemplate
	// If the given chain level does have a parent then get its key and certificate template.
	if chain.Parent != nil {
		parentKey, err = ReadKeyFile(chain.Parent.KeyFilename, chain.Parent.KeyPassword)
		if err != nil {
			return err
		}
		parentCertTemplate, err = ReadCertFile(chain.Parent.Filename)
		if err != nil {
			return err
		}
	}
	// Create a X.509v3 certificate.
	certBytes, err := x509.CreateCertificate(
		rand.Reader,
		certTemplate,
		parentCertTemplate,
		&key.PublicKey,
		parentKey,
	)
	if err != nil {
		return fmt.Errorf("failed to create certificate %v: %w", chain.Filename, err)
	}
	// Encode the certificate in a PEM block.
	pemBlock := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	}
	// Write PEM block to file.
	if _, err = file.Write(pem.EncodeToMemory(pemBlock)); err != nil {
		return fmt.Errorf("failed to write %v: %w", chain.Filename, err)
	}
	// Close file.
	if err := file.Close(); err != nil {
		return fmt.Errorf("failed to close %v: %w", chain.Filename, err)
	}
	return nil
}

// keyUsage returns the set of actions that are valid for a leaf or non-leaf key.
func keyUsage(leaf bool) x509.KeyUsage {
	if leaf {
		return x509.KeyUsageDigitalSignature
	}
	return x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign
}

// extKeyUsage returns the extended set of actions that are valid for a leaf or
// non-leaf key.
func extKeyUsage(leaf bool) []x509.ExtKeyUsage {
	if leaf {
		return []x509.ExtKeyUsage{
			x509.ExtKeyUsageClientAuth,
			x509.ExtKeyUsageServerAuth,
		}
	}
	return []x509.ExtKeyUsage{
		x509.ExtKeyUsageServerAuth,
	}
}
