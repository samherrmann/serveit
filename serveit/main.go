package main

import (
	"net/http"

	"github.com/samherrmann/serveit/handlers"
)

func main() {
	// Parse command-line flags into a configuration object.
	config := parseFlags()
	// Register file handler.
	http.HandleFunc("/", handlers.FileHandler(config.NotFoundFile))
	// Start HTTP server.
	listenAndServe(config.Port, config.TLS, config.Hosts)
}
