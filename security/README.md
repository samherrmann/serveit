# security

The security package provides a simple way to quickly generate self-signed
certificates.

## Example

```go
package main

import (
	"log"

	"github.com/samherrmann/serveit/security"
)

func main() {
	// 1. Define the chain of trust
	chain := &security.ChainOfTrust{
		Filename:    "my_app.crt",
		KeyFilename: "my_app.key",
		Subject: &security.Subject{
			CommonName:   "My Awesome App Name",
			Organization: []string{"My Awesome App Making Company Inc."},
			Country:      []string{"CA"},
			Province:     []string{"ON"},
			Locality:     []string{"Ottawa"},
		},
		Days:  3650,
		Hosts: []string{"localhost,192.168.0.1,example.com"},
		Parent: &security.ChainOfTrust{
			Filename:    "my_root_ca.crt",
			KeyFilename: "my_root_ca.key",
			KeyPassword: "mysecurepassword",
			Subject: &security.Subject{
				CommonName:   "My Awesome Certificate Authority",
				Organization: []string{"My Awesome Certificate Authority"},
				Country:      []string{"CA"},
				Province:     []string{"ON"},
				Locality:     []string{"Ottawa"},
			},
			Days: 3650,
		},
	}

	// 2. Create the RSA private key and x.509 certificate files.
	if err := security.WriteChainOfTrustFiles(chain, 0600); err != nil {
		log.Fatalln(err)
	}
}
```

The code above writes the following files to the working directory:
* `my_root_ca.key`
* `my_root_ca.crt`
* `my_app.key`
* `my_app.crt`
