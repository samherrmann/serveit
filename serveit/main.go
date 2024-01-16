package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/samherrmann/serveit/flag"
	"github.com/samherrmann/serveit/handlers"
)

func main() {
	if err := app(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func app() error {
	// Parse command-line flags into a configuration object.
	config, err := flag.Parse(os.Args)
	if err != nil {
		return err
	}
	// Register file handler.
	http.HandleFunc("/", handlers.FileHandler(config.NotFoundFile))
	// Start HTTP server.
	return listenAndServe(config.Port, config.TLS, config.Hosts)
}
