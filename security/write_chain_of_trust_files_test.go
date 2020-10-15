package security_test

import (
	"crypto/tls"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/samherrmann/serveit/security"
)

func TestWriteChainOfTrust(t *testing.T) {
	chain := newChainOfTrust("testdata", []string{"localhost,192.168.0.1"})

	// Start with a clean slate.
	if err := removeAllFiles(chain); err != nil {
		t.Error(err)
	}

	if err := security.WriteChainOfTrustFiles(chain); err != nil {
		t.Error(err)
	}

	if err := verifyKeyPair(chain); err != nil {
		t.Error(err)
	}

	// Clean up.
	if err := removeAllFiles(chain); err != nil {
		t.Error(err)
	}
}

// newChainOfTrust returns a chain of trust.
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

// verifyKeyPair returns nil when the public and private key match for all
// levels in the chain of trust.
func verifyKeyPair(chain *security.ChainOfTrust) error {
	return walkChain(chain, func(c *security.ChainOfTrust) error {
		_, err := tls.LoadX509KeyPair(c.Filename, c.KeyFilename)
		return err
	})
}

// removeAllFiles removes all key and certificate files in the chain.
func removeAllFiles(chain *security.ChainOfTrust) error {
	remove := func(cert *security.ChainOfTrust) error {
		rm := func(path string) error {
			if err := os.Remove(path); err != nil && !errors.Is(err, os.ErrNotExist) {
				return fmt.Errorf("cannot remove %v: %w", path, err)
			}
			return nil
		}
		if err := rm(cert.KeyFilename); err != nil {
			return err
		}
		if err := rm(cert.Filename); err != nil {
			return err
		}
		return nil
	}
	return walkChain(chain, remove)
}

// walkChain walks through the chain of trust
func walkChain(cert *security.ChainOfTrust, fn func(c *security.ChainOfTrust) error) error {
	if err := fn(cert); err != nil {
		return err
	}
	if cert.Parent != nil {
		walkChain(cert.Parent, fn)
	}
	return nil
}
