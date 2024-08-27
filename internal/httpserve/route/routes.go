package route

import (
	"time"

	echo "github.com/theopenlane/echox"
	"github.com/theopenlane/echox/middleware"

	"github.com/theopenlane/core/pkg/middleware/ratelimit"

	"github.com/theopenlane/dbx/internal/httpserve/handlers"
)

const (
	V1Version   = "v1"
	unversioned = ""
)

var (
	mw = []echo.MiddlewareFunc{middleware.Recover()}

	restrictedRateLimit = &ratelimit.Config{
		RateLimit:  1,
		BurstLimit: 1,
		ExpiresIn:  15 * time.Minute, // nolint:mnd
	}
	restrictedEndpointsMW = []echo.MiddlewareFunc{}
)

type Route struct {
	Method      string
	Path        string
	Handler     echo.HandlerFunc
	Middlewares []echo.MiddlewareFunc

	Name string
}

// RegisterRoutes with the echo routers
func RegisterRoutes(router *echo.Echo, h *handlers.Handler) error {
	// Middleware for restricted endpoints
	restrictedEndpointsMW = append(restrictedEndpointsMW, mw...)
	restrictedEndpointsMW = append(restrictedEndpointsMW, ratelimit.RateLimiterWithConfig(restrictedRateLimit)) // add restricted ratelimit middleware

	// routeHandlers that take the router and handler as input
	routeHandlers := []interface{}{
		// add handlers here
		registerReadinessHandler,
	}

	for _, route := range routeHandlers {
		if err := route.(func(*echo.Echo, *handlers.Handler) error)(router, h); err != nil {
			return err
		}
	}

	// register additional handlers that only require router input
	additionalHandlers := []interface{}{
		registerLivenessHandler,
		registerMetricsHandler,
	}

	for _, route := range additionalHandlers {
		if err := route.(func(*echo.Echo) error)(router); err != nil {
			return err
		}
	}

	return nil
}

// RegisterRoute with the echo server given a method, path, and handler definition
func (r *Route) RegisterRoute(router *echo.Echo) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method:      r.Method,
		Path:        r.Path,
		Handler:     r.Handler,
		Middlewares: r.Middlewares,

		Name: r.Name,
	})

	return
}
