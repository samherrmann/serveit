package security_test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/samherrmann/serveit/security"
)

func TestEnsureKeyPairs(t *testing.T) {
	// Start with a clean slate.
	if err := removeAllFiles(); err != nil {
		t.Error(err)
	}

	if err := security.EnsureKeyPairs([]string{"localhost"}); err != nil {
		t.Error(err)
	}

	if err := verifyKey(security.RootCAKeyFilename); err != nil {
		t.Error(err)
	}

	if err := verifyCert(security.RootCACertFilename); err != nil {
		t.Error(err)
	}

	if err := verifyKey(security.KeyFilename); err != nil {
		t.Error(err)
	}

	if err := verifyCert(security.CertFilename); err != nil {
		t.Error(err)
	}

	// Clean up.
	if err := removeAllFiles(); err != nil {
		t.Error(err)
	}
}

func verifyKey(filename string) error {
	return verifyFileContent(
		filename,
		"-----BEGIN RSA PRIVATE KEY-----",
		"-----END RSA PRIVATE KEY-----",
	)
}

func verifyCert(filename string) error {
	return verifyFileContent(
		filename,
		"-----BEGIN CERTIFICATE-----",
		"-----END CERTIFICATE-----",
	)
}

func verifyFileContent(filename string, prefix string, suffix string) error {
	// Read content from file.
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Error reading %v: %v", filename, err)
	}

	// Verify file content.
	contentStr := string(content)
	if !strings.HasPrefix(contentStr, prefix) &&
		!strings.HasSuffix(contentStr, suffix) {
		return fmt.Errorf("%v does not have expected content, got %v", filename, contentStr)
	}
	return nil
}

func removeAllFiles() error {
	files := []string{
		security.RootCAKeyFilename,
		security.RootCACertFilename,
		security.RootCACertSerialFilename,
		security.KeyFilename,
		security.CSRFilename,
		security.ExtFilename,
		security.CertFilename,
	}

	for _, file := range files {
		if err := os.Remove(file); err != nil && !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("Error removing %v: %w", file, err)
		}
	}
	return nil
}
