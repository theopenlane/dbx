package dbxclient

import (
	"net/http"

	"github.com/Yamashou/gqlgenc/clientv2"
)

type Config struct {
	// Enabled is a flag to enable the dbx client
	Enabled bool `json:"enabled" koanf:"enabled" jsonschema:"description=Enable the dbx client" default:"true"`
	// BaseURL is the base url for the dbx service
	BaseURL string `json:"baseUrl" koanf:"baseUrl" jsonschema:"description=Base URL for the dbx service" default:"http://localhost:1337"`
	// Endpoint is the endpoint for the graphql api
	Endpoint string `json:"endpoint" koanf:"endpoint" jsonschema:"description=Endpoint for the graphql api" default:"query"`
	// Debug is a flag to enable debug mode
	Debug bool `json:"debug" koanf:"debug" jsonschema:"description=Enable debug mode" default:"false"`
}

// NewDefaultClient creates a new default dbx client based on the config
func (c Config) NewDefaultClient() Dbxclient {
	i := WithEmptyInterceptor()
	interceptors := []clientv2.RequestInterceptor{i}

	if c.Debug {
		interceptors = append(interceptors, WithLoggingInterceptor())
	}

	return c.NewClientWithInterceptors(interceptors)
}

// NewClientWithInterceptors creates a new default dbx client with the provided interceptors
func (c Config) NewClientWithInterceptors(i []clientv2.RequestInterceptor) Dbxclient {
	h := http.DefaultClient

	// set options
	opts := &clientv2.Options{
		ParseDataAlongWithErrors: false,
	}

	gc := NewClient(h, c.BaseURL, opts, i...)

	return gc
}
