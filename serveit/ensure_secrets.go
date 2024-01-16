package main

import (
	"github.com/samherrmann/serveit/execpath"
	"github.com/samherrmann/serveit/security"
)

// KeyFilename is the filename of the private key.
type KeyFilename = string

// CertFilename is the filename of the certificate.
type CertFilename = string

func ensureSecrets(hosts []string) (KeyFilename, CertFilename, error) {
	dir, err := execpath.Dir()
	if err != nil {
		return "", "", err
	}
	chain := newChainOfTrust(dir, hosts)
	if err := security.WriteChainOfTrustFiles(chain, 0600); err != nil {
		return "", "", err
	}
	return chain.KeyFilename, chain.Filename, nil
}
