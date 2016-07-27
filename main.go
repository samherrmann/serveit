package main

import (
	"flag"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
)

func main() {

	port := flag.Int("port", 8080, "The port on which to serve the current directory.")
	flag.Parse()

	addr := ":" + strconv.Itoa(*port)
	path, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}

	fmt.Println("Serving root \"" + path + "\" on port " + addr)
	panic(http.ListenAndServe(addr, http.FileServer(http.Dir(path))))
}
