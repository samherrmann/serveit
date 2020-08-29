package main

import (
	"flag"
	"os"
)

// parseFlags returns a Config object containing the values of the commnad-line flags.
func parseFlags(args []string) *Config {
	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	port := flags.Int(
		"port",
		8080,
		"The port on which to serve the current directory.",
	)
	spaMode := flags.Bool(
		"spa",
		false,
		"Single-page application mode: If set, requests for which no file or directory exists are redirected to index.html.",
	)
	flags.Parse(args)
	return &Config{
		Port:    *port,
		SPAMode: *spaMode,
	}
}
