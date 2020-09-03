package handlers

import (
	"net/http"
	"os"
)

// FileHandler returns a function to handle file requests. If the provided
// notFoundFile argument is a valid path to a file, then that file is served
// when the requested resource is not found. Set notFoundFile to an empty string
// to serve Go's default "404 page not found" response when the requested
// resource is not found.
func FileHandler(notFoundFile string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !doesPathExist(r) && notFoundFile != "" {
			w.WriteHeader(http.StatusNotFound)
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
