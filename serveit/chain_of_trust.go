package main

import (
	"path/filepath"

	"github.com/samherrmann/serveit/security"
)

func newChainOfTrust(dir string, hosts []string) *security.ChainOfTrust {
	return &security.ChainOfTrust{
		Filename:    filepath.Join(dir, "serveit.crt"),
		KeyFilename: filepath.Join(dir, "serveit.key"),
		Subject: &security.Subject{
			CommonName:   "Serveit",
			Organization: []string{"github.com/samherrmann"},
			Country:      []string{"CA"},
			Province:     []string{"ON"},
			Locality:     []string{"Ottawa"},
		},
		Days:  3650,
		Hosts: hosts,
		Parent: &security.ChainOfTrust{
			Filename:    filepath.Join(dir, "serveit_root_ca.crt"),
			KeyFilename: filepath.Join(dir, "serveit_root_ca.key"),
			KeyPassword: "serveit",
			Subject: &security.Subject{
				CommonName:   "Serveit Root Certificate Authority",
				Organization: []string{"github.com/samherrmann"},
				Country:      []string{"CA"},
				Province:     []string{"ON"},
				Locality:     []string{"Ottawa"},
			},
			Days: 3650,
		},
	}
}
