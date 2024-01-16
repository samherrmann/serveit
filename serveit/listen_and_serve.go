package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func listenAndServe(port int, tls bool, hosts []string) error {
	addr := ":" + strconv.Itoa(port)
	fmt.Println("Serving current directory on port " + addr)

	if tls {
		keyFilename, certFilename, err := ensureSecrets(hosts)
		if err != nil {
			return err
		}
		return http.ListenAndServeTLS(addr, certFilename, keyFilename, nil)
	}

	return http.ListenAndServe(addr, nil)
}
