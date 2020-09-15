package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

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
	listenAndServe(config.Port, config.TLS)
}

func parseFlags() *flag.Config {
	config, err := flag.Parse(os.Args[1:])
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}
	return config
}

func ensureSecrets() {
	if err := security.EnsureKeyPairs(); err != nil {
		log.Fatalln(err)
	}
}

func listenAndServe(port int, tls bool) {
	addr := ":" + strconv.Itoa(port)
	log.Println("Serving current directory on port " + addr)
	var err error
	if tls {
		ensureSecrets()
		err = http.ListenAndServeTLS(addr, security.CertFilename, security.KeyFilename, nil)
	} else {
		err = http.ListenAndServe(addr, nil)
	}
	if err != nil {
		log.Fatalln(err)
	}
}
