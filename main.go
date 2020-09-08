package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/samherrmann/serveit/flag"
	"github.com/samherrmann/serveit/handlers"
)

func main() {
	// Parse command-line flags into a configuration object
	config := parseFlags()
	// Register file handler
	http.HandleFunc("/", handlers.FileHandler(config.NotFoundFile))
	// Start HTTP server
	listenAndServe(config.Port)
}

func parseFlags() *flag.Config {
	config, err := flag.Parse(os.Args[1:])
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}
	return config
}

func listenAndServe(port int) {
	addr := ":" + strconv.Itoa(port)
	log.Println("Serving current directory on port " + addr)
	log.Fatalln(http.ListenAndServe(addr, nil))
}
