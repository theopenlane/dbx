package graphapi

import (
	"fmt"
	"net/http"
	"time"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"
	"github.com/ravilushqa/otelgqlgen"
	echo "github.com/theopenlane/echox"
	"github.com/wundergraph/graphql-go-tools/pkg/playground"
	"go.uber.org/zap"

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
	graphPath      = "query"
	playgroundPath = "playground"

	graphFullPath = fmt.Sprintf("/%s", graphPath)
)

// Resolver provides a graph response resolver
type Resolver struct {
	client *ent.Client
	logger *zap.SugaredLogger
}

// NewResolver returns a resolver configured with the given ent client
func NewResolver(client *ent.Client) *Resolver {
	return &Resolver{
		client: client,
	}
}

func (r Resolver) WithLogger(l *zap.SugaredLogger) *Resolver {
	r.logger = l

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
	srv := handler.NewDefaultServer(
		NewExecutableSchema(
			Config{
				Resolvers: r,
			},
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

	srv.SetQueryCache(lru.New(1000)) // nolint:mnd

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100), // nolint:mnd
	})
	// add transactional db client
	WithTransactions(srv, r.client)

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

// WithTransactions adds the transactioner to the ent db client
func WithTransactions(h *handler.Server, c *ent.Client) {
	// setup transactional db client
	h.AroundOperations(injectClient(c))
	h.Use(entgql.Transactioner{TxOpener: c})
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
			h.r.logger.Fatal("error configuring playground handlers", "error", err)
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
