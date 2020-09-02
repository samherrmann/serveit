package flag

// Config contains the application configuration.
type Config struct {
	// Port number to listen on.
	Port int
	// The path of the file to serve when the requested resource cannot be found.
	NotFoundFile string
}
