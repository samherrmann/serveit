package main

import (
	"log"

	"github.com/samherrmann/serveit/execpath"
	"github.com/samherrmann/serveit/security"
)

// KeyFilename is the filename of the private key.
type KeyFilename = string

// CertFilename is the filename of the certificate.
type CertFilename = string

func ensureSecrets(hosts []string) (KeyFilename, CertFilename) {
	dir, err := execpath.Dir()
	if err != nil {
		log.Fatalln(err)
	}
	chain := newChainOfTrust(dir, hosts)
	if err := security.WriteChainOfTrustFiles(chain, 0600); err != nil {
		log.Fatalln(err)
	}
	return chain.KeyFilename, chain.Filename
}
