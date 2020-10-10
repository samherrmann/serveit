package security

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

// ReadKeyFile returns the RSA private key from the file named by the provided
// filename. If the password is not an empty string then it's used to decrypt
// the PEM block.
func ReadKeyFile(filename, password string) (*rsa.PrivateKey, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("cannot read %v: %w", filename, err)
	}
	block, _ := pem.Decode(bytes)

	var blockBytes []byte
	if password != "" {
		blockBytes, err = x509.DecryptPEMBlock(block, []byte(password))
		if err != nil {
			return nil, fmt.Errorf("cannot decrypt %v: %w", filename, err)
		}
	} else {
		blockBytes = block.Bytes
	}
	key, err := x509.ParsePKCS1PrivateKey(blockBytes)
	if err != nil {
		return nil, fmt.Errorf("cannot parse %v: %w", filename, err)
	}
	if err := key.Validate(); err != nil {
		return nil, fmt.Errorf("key %v is invalid: %w", filename, err)
	}
	return key, nil
}
