package server

// Server define a server behaviour
type Server interface {
	// Serve serves a service's server implementation
	Serve() error
}

// Config contains the configuration to set up the server
type Config struct {
	Protocol string
	Host     string
	Port     string
}

