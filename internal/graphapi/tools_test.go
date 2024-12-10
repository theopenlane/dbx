package graphapi_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/Yamashou/gqlgenc/clientv2"
	"github.com/theopenlane/go-turso"
	"github.com/theopenlane/utils/testutils"

	ent "github.com/theopenlane/dbx/internal/ent/generated"
	"github.com/theopenlane/dbx/internal/entdb"
	"github.com/theopenlane/dbx/internal/graphapi"
	"github.com/theopenlane/dbx/pkg/dbxclient"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

// TestGraphTestSuite runs all the tests in the GraphTestSuite
func TestGraphTestSuite(t *testing.T) {
	suite.Run(t, new(GraphTestSuite))
}

// GraphTestSuite handles the setup and teardown between tests
type GraphTestSuite struct {
	suite.Suite
	client *client
	tf     *testutils.TestFixture
}

// client contains all the clients the test need to interact with
type client struct {
	db  *ent.Client
	dbx dbxclient.Dbxclient
}

type graphClient struct {
	srvURL     string
	httpClient *http.Client
}

func (suite *GraphTestSuite) SetupSuite() {
	zerolog.SetGlobalLevel(zerolog.Disabled)

	suite.tf = entdb.NewTestFixture()
}

func (suite *GraphTestSuite) SetupTest() {
	t := suite.T()

	ctx := context.Background()

	// setup mock turso client
	tc := turso.NewMockClient()

	opts := []ent.Option{
		ent.Turso(tc),
	}

	// create database connection
	db, err := entdb.NewTestClient(ctx, suite.tf, opts)
	if err != nil {
		require.NoError(t, err, "failed opening connection to database")
	}

	// assign values
	c := &client{
		db:  db,
		dbx: graphTestClient(t, db),
	}

	suite.client = c
}

func (suite *GraphTestSuite) TearDownTest() {
	if suite.client.db != nil {
		if err := suite.client.db.Close(); err != nil {
			log.Fatal().Err(err).Msg("failed to close database")
		}
	}
}

func (suite *GraphTestSuite) TearDownSuite() {
	testutils.TeardownFixture(suite.tf)
}

func graphTestClient(t *testing.T, c *ent.Client) dbxclient.Dbxclient {
	srv := handler.New(
		graphapi.NewExecutableSchema(
			graphapi.Config{Resolvers: graphapi.NewResolver(c)},
		))

	// add all the transports to the server
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	graphapi.WithTransactions(srv, c)

	g := &graphClient{
		srvURL:     "query",
		httpClient: &http.Client{Transport: localRoundTripper{handler: srv}},
	}

	// set options
	opt := &clientv2.Options{
		ParseDataAlongWithErrors: false,
	}

	// setup interceptors
	i := dbxclient.WithEmptyInterceptor()

	return dbxclient.NewClient(g.httpClient, g.srvURL, opt, i)
}

// localRoundTripper is an http.RoundTripper that executes HTTP transactions
// by using handler directly, instead of going over an HTTP connection.
type localRoundTripper struct {
	handler http.Handler
}

func (l localRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	w := httptest.NewRecorder()
	l.handler.ServeHTTP(w, req)

	return w.Result(), nil
}
