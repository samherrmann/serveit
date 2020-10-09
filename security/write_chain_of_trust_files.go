package security

import (
	"errors"
	"os"
)

// WriteChainOfTrustFiles writes all RSA private keys and x.509 certificates
// defined in ChainOfTrust. If the files defined in the chain of trust
// already exist they are not overwritten and no error is returned.
func WriteChainOfTrustFiles(chain *ChainOfTrust) error {
	chain.leaf = true
	invertedChain := []ChainOfTrust{}
	var createChain func(*ChainOfTrust)

	createChain = func(cert *ChainOfTrust) {
		invertedChain = append([]ChainOfTrust{*cert}, invertedChain...)
		if cert.Parent != nil {
			createChain(cert.Parent)
		}
	}
	createChain(chain)

	// isError returns true for any error except file-exists error.
	isError := func(err error) bool {
		return err != nil && !errors.Is(err, os.ErrExist)
	}

	for _, c := range invertedChain {
		if err := WriteKeyFile(c.KeyFilename, c.KeyPassword); isError(err) {
			return err
		}
		if err := WriteCertFile(&c); isError(err) {
			return err
		}
	}
	return nil
}
