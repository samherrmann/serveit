package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var (
	port               *int
	isAngularRoutingOn *bool
)

func main() {
	parseFlags()
	http.HandleFunc("/", serveFile)
	listenAndServe()
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	if !doesPathExist(r) && *isAngularRoutingOn {
		http.ServeFile(w, r, ".")
		return
	}
	http.ServeFile(w, r, r.URL.Path[1:])
}

func parseFlags() {
	port = flag.Int("port", 8080, "The port on which to serve the current directory.")
	isAngularRoutingOn = flag.Bool("ar", false, "Angular routing: If set, requests for which no file or directory exists are redirected to the root.")
	flag.Parse()
}

func doesPathExist(r *http.Request) bool {
	_, err := os.Stat("." + r.URL.Path)
	return err == nil
}

func listenAndServe() {
	addr := ":" + strconv.Itoa(*port)
	fmt.Println("Serving current directory on port " + addr)
	panic(http.ListenAndServe(addr, nil))
}
