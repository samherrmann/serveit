package flag

// Config contains the application configuration.
type Config struct {
	// Port number to listen on.
	Port int
	// The path of the file to serve when the requested resource cannot be found.
	NotFoundFile string
	// When true the app will use HTTPS instead of HTTP.
	TLS bool
	// Hostnames to add to the server X.509 certificate.
	Hostnames []string
}
