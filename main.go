package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/samherrmann/serveit/flag"
	"github.com/samherrmann/serveit/handlers"
)

func main() {
	// Parse command-line flags into a configuration object
	config := flag.Parse()
	// Register file handler
	http.HandleFunc("/", handlers.FileHandler(config.NotFoundFile))
	// Start HTTP server
	listenAndServe(config.Port)
}

func listenAndServe(port int) {
	addr := ":" + strconv.Itoa(port)
	log.Println("Serving current directory on port " + addr)
	log.Fatalln(http.ListenAndServe(addr, nil))
}
