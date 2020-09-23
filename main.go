package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/samherrmann/serveit/execpath"
	"github.com/samherrmann/serveit/flag"
	"github.com/samherrmann/serveit/handlers"
	"github.com/samherrmann/serveit/security"
)

func main() {
	// Parse command-line flags into a configuration object.
	config := parseFlags()
	// Register file handler.
	http.HandleFunc("/", handlers.FileHandler(config.NotFoundFile))
	// Start HTTP server.
	listenAndServe(config.Port, config.TLS, config.Hostnames)
}

func parseFlags() *flag.Config {
	config, err := flag.Parse(os.Args[1:])
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}
	return config
}

func ensureSecrets(dir string, hostnames []string) {
	if err := security.EnsureKeyPairs(dir, hostnames); err != nil {
		log.Fatalln(err)
	}
}

func listenAndServe(port int, tls bool, hostnames []string) {
	addr := ":" + strconv.Itoa(port)
	log.Println("Serving current directory on port " + addr)
	if tls {
		dir, err := execpath.Dir()
		if err != nil {
			log.Fatalln(err)
		}
		ensureSecrets(dir, hostnames)
		keyPath := filepath.Join(dir, security.KeyFilename)
		certPath := filepath.Join(dir, security.CertFilename)
		log.Fatalln(http.ListenAndServeTLS(addr, certPath, keyPath, nil))
	} else {
		log.Fatalln(http.ListenAndServe(addr, nil))
	}
}
