package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/samherrmann/serveit/flag"
)

func main() {
	// Parse command-line flags into a configuration object
	config := flag.Parse()
	// Register file handler
	http.HandleFunc("/", fileHandler(config.NotFoundFile))
	// Start HTTP server
	listenAndServe(config.Port)
}

func fileHandler(notFoundFile string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !doesPathExist(r) && notFoundFile != "" {
			http.ServeFile(w, r, notFoundFile)
			return
		}
		http.ServeFile(w, r, r.URL.Path[1:])
	}
}

func doesPathExist(r *http.Request) bool {
	_, err := os.Stat("." + r.URL.Path)
	return err == nil
}

func listenAndServe(port int) {
	addr := ":" + strconv.Itoa(port)
	fmt.Println("Serving current directory on port " + addr)
	panic(http.ListenAndServe(addr, nil))
}
