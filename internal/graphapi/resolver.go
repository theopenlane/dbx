package graphapi

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"ariga.io/entcache"
	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"
	"github.com/ravilushqa/otelgqlgen"
	"github.com/rs/zerolog/log"
	echo "github.com/theopenlane/echox"
	"github.com/theopenlane/gqlgen-plugins/graphutils"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/wundergraph/graphql-go-tools/pkg/playground"

	ent "github.com/theopenlane/dbx/internal/ent/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

const (
	ActionGet    = "get"
	ActionUpdate = "update"
	ActionDelete = "delete"
	ActionCreate = "create"
)

var (
	graphPath               = "query"
	playgroundPath          = "playground"
	defaultComplexityLimit  = 100
	introspectionComplexity = 200

	graphFullPath = fmt.Sprintf("/%s", graphPath)
)

// Resolver provides a graph response resolver
type Resolver struct {
	client            *ent.Client
	extensionsEnabled bool
	isDevelopment     bool
	complexityLimit   int
	maxResultLimit    *int
}

// NewResolver returns a resolver configured with the given ent client
func NewResolver(client *ent.Client) *Resolver {
	return &Resolver{
		client: client,
	}
}

func (r Resolver) WithExtensions(enabled bool) *Resolver {
	r.extensionsEnabled = enabled

	return &r
}

// WithDevelopment sets the resolver to development mode
// when isDevelopment is false, introspection will be disabled
func (r Resolver) WithDevelopment(dev bool) *Resolver {
	r.isDevelopment = dev

	return &r
}

// WithComplexityLimitConfig sets the complexity limit for the resolver
func (r Resolver) WithComplexityLimitConfig(limit int) *Resolver {
	r.complexityLimit = limit

	return &r
}

// WithMaxResultLimit sets the max result limit in the config for the resolvers
func (r Resolver) WithMaxResultLimit(limit int) *Resolver {
	r.maxResultLimit = &limit

	return &r
}

// Handler is an http handler wrapping a Resolver
type Handler struct {
	r              *Resolver
	graphqlHandler *handler.Server
	playground     *playground.Playground
	middleware     []echo.MiddlewareFunc
}

// Handler returns an http handler for a graph resolver
func (r *Resolver) Handler(withPlayground bool) *Handler {
	c := Config{Resolvers: r}

	srv := handler.New(
		NewExecutableSchema(
			c,
		),
	)

	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second, // nolint:mnd
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000)) // nolint:mnd

	if r.isDevelopment {
		srv.Use(extension.Introspection{})
	}

	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100), // nolint:mnd
	})

	// add complexity limit
	r.WithComplexityLimit(srv)
	// add transactional db client
	WithTransactions(srv, r.client)

	// add max result limits to fields in requests
	WithResultLimit(srv, r.maxResultLimit)

	srv.Use(otelgqlgen.Middleware())

	h := &Handler{
		r:              r,
		graphqlHandler: srv,
	}

	if withPlayground {
		h.playground = playground.New(playground.Config{
			PathPrefix:          "/",
			PlaygroundPath:      playgroundPath,
			GraphqlEndpointPath: graphFullPath,
		})
	}

	return h
}

// WithComplexityLimit adds a complexity limit to the GraphQL handler
func (r *Resolver) WithComplexityLimit(h *handler.Server) {
	// prevent complex queries except the introspection query
	h.Use(&extension.ComplexityLimit{
		Func: func(_ context.Context, rc *graphql.OperationContext) int {
			if rc != nil && rc.OperationName == "IntrospectionQuery" {
				return introspectionComplexity
			}

			if rc.OperationName == "GlobalSearch" {
				// allow more complexity for the global search
				// e.g. if the complexity limit is 100, we allow 500 for the global search
				return r.complexityLimit * 5 //nolint:mnd
			}

			if r.complexityLimit > 0 {
				return r.complexityLimit
			}

			return defaultComplexityLimit
		},
	})
}

// WithTransactions adds the transactioner to the ent db client
func WithTransactions(h *handler.Server, c *ent.Client) {
	// setup transactional db client
	h.AroundOperations(injectClient(c))

	h.Use(entgql.Transactioner{TxOpener: c})
}

// WithContextLevelCache adds a context level cache to the handler
func WithContextLevelCache(h *handler.Server) {
	h.AroundResponses(func(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
		if op := graphql.GetOperationContext(ctx).Operation; op != nil && op.Operation == ast.Query {
			ctx = entcache.NewContext(ctx)
		}

		return next(ctx)
	})
}

// WithResultLimit adds a max result limit to the handler in order to set limits on
// all nested edges in the graphql request
func WithResultLimit(h *handler.Server, limit *int) {
	h.AroundFields(func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
		if limit == nil {
			return next(ctx)
		}

		// grab preloads to set max result limits
		graphutils.GetPreloads(ctx, limit)

		return next(ctx)
	})
}

// WithSkipCache adds a skip cache middleware to the handler
// This is useful for testing, where you don't want to cache responses
// so you can see the changes immediately
func WithSkipCache(h *handler.Server) {
	h.AroundResponses(func(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
		return next(entcache.Skip(ctx))
	})
}

// Handler returns the http.HandlerFunc for the GraphAPI
func (h *Handler) Handler() http.HandlerFunc {
	return h.graphqlHandler.ServeHTTP
}

// Routes for the the server
func (h *Handler) Routes(e *echo.Group) {
	e.Use(h.middleware...)

	// Create the default POST graph endpoint
	e.POST("/"+graphPath, func(c echo.Context) error {
		h.graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	// Create a GET query endpoint in order to create short queries with a query string
	e.GET("/"+graphPath, func(c echo.Context) error {
		h.graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	if h.playground != nil {
		handlers, err := h.playground.Handlers()
		if err != nil {
			log.Fatal().Err(err).Msg("error configuring playground handlers")
			return
		}

		for i := range handlers {
			// with the function we need to dereference the handler so that it remains
			// the same in the function below
			hCopy := handlers[i].Handler

			e.GET(handlers[i].Path, func(c echo.Context) error {
				hCopy.ServeHTTP(c.Response(), c.Request())
				return nil
			})
		}
	}
}
