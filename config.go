package main

// Config contains the application configuration.
type Config struct {
	// Port number to listen on.
	Port int
	// Enable single-page application mode if true. Disable otherwise.
	SPAMode bool
}
