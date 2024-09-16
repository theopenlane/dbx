package handlers

import (
	"github.com/redis/go-redis/v9"
	echo "github.com/theopenlane/echox"

	"github.com/theopenlane/iam/sessions"

	ent "github.com/theopenlane/dbx/internal/ent/generated"
)

// Handler contains configuration options for handlers
type Handler struct {
	// IsTest is a flag to determine if the application is running in test mode and will mock external calls
	IsTest bool
	// DBClient to interact with the generated ent schema
	DBClient *ent.Client
	// RedisClient to interact with redis
	RedisClient *redis.Client
	// ReadyChecks is a set of checkFuncs to determine if the application is "ready" upon startup
	ReadyChecks Checks
	// SessionConfig to handle sessions
	SessionConfig *sessions.SessionConfig
	// AuthMiddleware contains the middleware to be used for authenticated endpoints
	AuthMiddleware []echo.MiddlewareFunc
}
