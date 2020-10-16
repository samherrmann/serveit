package security

import "crypto/x509/pkix"

// Subject represents a x.509 certificate subject.
type Subject = pkix.Name

// ChainOfTrust defines a chain of x.509 certificates and RSA private keys.
type ChainOfTrust struct {
	// Parent is the parent link in the chain of trust.
	Parent *ChainOfTrust
	// Days is the number of days for which the certificate is valid.
	Days int
	// Subject is x.509 certificate subject.
	Subject *Subject
	// Hosts is a list of domain names and/or IP addresses.
	Hosts []string
	// Filename is the filename of the file containing the x.509 certificate.
	Filename string
	// KeyFilename is the filename of the file containing the RSA private key.
	KeyFilename string
	// KeyPassword is the password used to encrypt the RSA private key PEM block.
	// If no password is provided (i.e. empty string) then the PEM block is not
	// encrypted.
	KeyPassword string
	// leaf defines wether this level of the chain of trust is the leaf.
	leaf bool
}
