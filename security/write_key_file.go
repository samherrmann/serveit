package security

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// WriteKeyFile creates a RSA private key and writes it to the file named by the
// provided filename. The key is written as a PEM block. If the password is not
// an empty string then the PEM block is encrypted using the password. Note that
// the file named by filename must not exist or an error is returned. A
// file-exist error may be checked with errors.Is(err, os.ErrExist).
func WriteKeyFile(filename, password string) error {
	// Attempt to create and open file. Error if file already exists.
	fileFlag := os.O_WRONLY | os.O_CREATE | os.O_EXCL
	file, err := os.OpenFile(filename, fileFlag, 0644)
	if err != nil {
		return fmt.Errorf("failed to open %v: %w", filename, err)
	}
	// Generate a 2048 bit RSA key.
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return fmt.Errorf("failed to generate RSA private key %v: %w", filename, err)
	}
	// Encode private key in a PEM block.
	pemBlock, err := newPrivateKeyPEMBlock(key, password)
	if err != nil {
		return fmt.Errorf("failed to generate PEM block for %v: %w", filename, err)
	}
	// Write PEM block to file.
	if _, err = file.Write(pem.EncodeToMemory(pemBlock)); err != nil {
		return fmt.Errorf("failed to write %v: %w", filename, err)
	}
	// Close file.
	if err := file.Close(); err != nil {
		return fmt.Errorf("failed to close %v: %w", filename, err)
	}
	return nil
}

func newPrivateKeyPEMBlock(key *rsa.PrivateKey, password string) (*pem.Block, error) {
	if password == "" {
		return &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		}, nil
	}
	return x509.EncryptPEMBlock(
		rand.Reader,
		"RSA PRIVATE KEY",
		x509.MarshalPKCS1PrivateKey(key),
		[]byte(password),
		x509.PEMCipherAES256,
	)
}
