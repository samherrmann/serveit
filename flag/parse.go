package flag

import (
	"flag"
	"os"
)

// Parse returns a Config object containing the values of the commnad-line
// flags. If an error occurs while parsing the command-line flags, then the
// program will exit with status code 2.
func Parse() *Config {
	config, err := parse(os.Args[1:])
	if err != nil {
		os.Exit(2)
	}
	return config
}

func parse(args []string) (*Config, error) {
	// A custom flag set is used instead of the static one built into Go's flag
	// package in order to make this function testable.
	flagSet := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
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
	if err := flagSet.Parse(args); err != nil {
		return nil, err
	}
	return &Config{
		Port:         *port,
		NotFoundFile: *notFoundFile,
	}, nil
}
