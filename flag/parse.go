package flag

import (
	"flag"
	"strings"
)

// Parse parses flag definitions from the argument list, where the first element
// is expected to be the program name. A Config object is returned containing
// the flag values. The pgiven argument list usually originates from os.Args.
func Parse(args []string) (*Config, error) {
	programName := ""
	if len(args) > 0 {
		programName = args[0]
	}
	if len(args) > 1 {
		args = args[1:]
	} else {
		args = []string{}
	}
	// A custom flag set is used instead of the static one built into Go's flag
	// package in order to make this function testable.
	flagSet := flag.NewFlagSet(programName, flag.ContinueOnError)
	port := flagSet.Int(
		"port",
		8080,
		"The port on which to serve the current directory.",
	)
	notFoundFile := flagSet.String(
		"not-found-file",
		"",
		"The path of the file to serve when the requested resource cannot be found. "+
			"For single-page applications, this flag is typically set to index.html.",
	)
	tls := flagSet.Bool(
		"tls",
		false,
		"When true, servit automatically generates a self-signed certificate and serves "+
			"files over HTTPS. Requires OpenSSL to be available on the system PATH.",
	)
	hosts := flagSet.String(
		"hosts",
		"localhost",
		"A comma-separated list (no spaces) of DNS names and/or IP addresses to add to the "+
			"server X.509 certificate. This flag is only applicable when the -tls flag is set.",
	)

	if err := flagSet.Parse(args); err != nil {
		return nil, err
	}

	return &Config{
		Port:         *port,
		NotFoundFile: *notFoundFile,
		TLS:          *tls,
		Hosts:        strings.Split(*hosts, ","),
	}, nil
}
