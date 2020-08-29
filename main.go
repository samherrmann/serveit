package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// Get command-line arguments without program name
	args := os.Args[1:]
	// Parse command-line flags into a configuration object
	config := parseFlags(args)
	// Register file handler
	http.HandleFunc("/", fileHandler(config.SPAMode))
	// Start HTTP server
	listenAndServe(config.Port)
}

func fileHandler(spaMode bool) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !doesPathExist(r) && spaMode {
			http.ServeFile(w, r, ".")
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
