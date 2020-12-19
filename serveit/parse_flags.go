package main

import (
	"log"
	"os"

	"github.com/samherrmann/serveit/flag"
)

func parseFlags() *flag.Config {
	config, err := flag.Parse(os.Args)
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}
	return config
}
