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

func TestEnsureKeyPairs(t *testing.T) {
	dir := "testdata"

	// Start with a clean slate.
	if err := removeAllFiles(dir); err != nil {
		t.Error(err)
	}

	if err := security.EnsureKeyPairs(dir, []string{"localhost"}); err != nil {
		t.Error(err)
	}

	if err := verifyKey(dir, security.RootCAKeyFilename); err != nil {
		t.Error(err)
	}

	if err := verifyCert(dir, security.RootCACertFilename); err != nil {
		t.Error(err)
	}

	if err := verifyKey(dir, security.KeyFilename); err != nil {
		t.Error(err)
	}

	if err := verifyCert(dir, security.CertFilename); err != nil {
		t.Error(err)
	}

	// Clean up.
	if err := removeAllFiles(dir); err != nil {
		t.Error(err)
	}
}

func verifyKey(dir, filename string) error {
	return verifyFileContent(
		dir,
		filename,
		"-----BEGIN RSA PRIVATE KEY-----",
		"-----END RSA PRIVATE KEY-----",
	)
}

func verifyCert(dir, filename string) error {
	return verifyFileContent(
		dir,
		filename,
		"-----BEGIN CERTIFICATE-----",
		"-----END CERTIFICATE-----",
	)
}

func verifyFileContent(dir, filename, prefix, suffix string) error {
	// Read content from file.
	content, err := ioutil.ReadFile(filepath.Join(dir, filename))
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

func removeAllFiles(dir string) error {
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
		if err := os.Remove(filepath.Join(dir, file)); err != nil && !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("Error removing %v: %w", file, err)
		}
	}
	return nil
}
