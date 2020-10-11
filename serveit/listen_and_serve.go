package main

import (
	"log"
	"net/http"
	"strconv"
)

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
