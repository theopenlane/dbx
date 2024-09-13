package server

import (
	"context"

	"github.com/rs/zerolog/log"
	echo "github.com/theopenlane/echox"

	echodebug "github.com/theopenlane/core/pkg/middleware/debug"

	"github.com/theopenlane/dbx/internal/httpserve/config"
	"github.com/theopenlane/dbx/internal/httpserve/route"
)

type Server struct {
	// config contains the base server settings
	config config.Config
	// handlers contains additional handlers to register with the echo server
	handlers []handler
}

type handler interface {
	Routes(*echo.Group)
}

// AddHandler provides the ability to add additional HTTP handlers that process
// requests. The handler that is provided should have a Routes(*echo.Group)
// function, which allows the routes to be added to the server.
func (s *Server) AddHandler(r handler) {
	s.handlers = append(s.handlers, r)
}

// NewServer returns a new Server configuration
func NewServer(c config.Config) *Server {
	return &Server{
		config: c,
	}
}

// StartEchoServer creates and starts the echo server with configured middleware and handlers
func (s *Server) StartEchoServer(ctx context.Context) error {
	srv := echo.New()

	sc := echo.StartConfig{
		HideBanner:      true,
		HidePort:        true,
		Address:         s.config.Settings.Server.Listen,
		GracefulTimeout: s.config.Settings.Server.ShutdownGracePeriod,
		GracefulContext: ctx,
	}

	srv.Debug = s.config.Settings.Server.Debug

	if srv.Debug {
		srv.Use(echodebug.BodyDump())
	}

	for _, m := range s.config.DefaultMiddleware {
		srv.Use(m)
	}

	// Add base routes to the server
	if err := route.RegisterRoutes(srv, &s.config.Handler); err != nil {
		return err
	}

	// Registers additional routes for the graph endpoints with middleware defined
	for _, handler := range s.handlers {
		handler.Routes(srv.Group("", s.config.GraphMiddleware...))
	}

	// Print routes on startup
	routes := srv.Router().Routes()
	for _, r := range routes {
		log.Info().Str("route", r.Path()).Str("method", r.Method()).Msg("registered route")
	}

	// if TLS is enabled, start new echo server with TLS
	if s.config.Settings.Server.TLS.Enabled {
		log.Info().Msg("starting in https mode")

		return sc.StartTLS(srv, s.config.Settings.Server.TLS.CertFile, s.config.Settings.Server.TLS.CertKey)
	}

	// otherwise, start without TLS
	return sc.Start(srv)
}
