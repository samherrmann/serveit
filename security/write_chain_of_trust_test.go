package security_test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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

	if err := verifyKeys(chain); err != nil {
		t.Error(err)
	}

	if err := verifyCerts(chain); err != nil {
		t.Error(err)
	}

	// Clean up.
	if err := removeAllFiles(chain); err != nil {
		t.Error(err)
	}
}

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

func verifyKeys(chain *security.ChainOfTrust) error {
	return walkChain(chain, func(c *security.ChainOfTrust) error {
		return verifyFileContent(
			c.KeyFilename,
			"-----BEGIN RSA PRIVATE KEY-----",
			"-----END RSA PRIVATE KEY-----",
		)
	})
}

func verifyCerts(chain *security.ChainOfTrust) error {
	return walkChain(chain, func(c *security.ChainOfTrust) error {
		return verifyFileContent(
			c.Filename,
			"-----BEGIN CERTIFICATE-----",
			"-----END CERTIFICATE-----",
		)
	})
}

// verifyFileContent verifies that the file content start and end with the
// provided prefix and postfix respectively.
func verifyFileContent(filename, prefix, suffix string) error {
	// Read content from file.
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read %v: %w", filename, err)
	}

	// Verify file content.
	contentStr := string(content)
	if !strings.HasPrefix(contentStr, prefix) &&
		!strings.HasSuffix(contentStr, suffix) {
		return fmt.Errorf("%v does not have expected content, got %v", filename, contentStr)
	}
	return nil
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

func walkChain(cert *security.ChainOfTrust, fn func(c *security.ChainOfTrust) error) error {
	if err := fn(cert); err != nil {
		return err
	}
	if cert.Parent != nil {
		walkChain(cert.Parent, fn)
	}
	return nil
}
