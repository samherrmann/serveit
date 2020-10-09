package main

import (
	"log"
	"net/http"
	"strconv"

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

func listenAndServe(port int, tls bool, hosts []string) {
	addr := ":" + strconv.Itoa(port)
	log.Println("Serving current directory on port " + addr)
	if tls {
		keyFilename, certFilename := ensureSecrets(hosts)
		log.Fatalln(http.ListenAndServeTLS(addr, certFilename, keyFilename, nil))
	} else {
		log.Fatalln(http.ListenAndServe(addr, nil))
	}
}
