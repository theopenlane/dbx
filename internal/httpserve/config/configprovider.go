package config

// Provider serves as a common interface to read echo server configuration
type Provider interface {
	// GetConfig returns the server configuration
	GetConfig() (*Config, error)
}
