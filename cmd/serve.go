package cmd

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theopenlane/beacon/otelx"
	"github.com/theopenlane/go-turso"
	"github.com/theopenlane/utils/cache"

	ent "github.com/theopenlane/dbx/internal/ent/generated"
	"github.com/theopenlane/dbx/internal/entdb"
	"github.com/theopenlane/dbx/internal/httpserve/config"
	"github.com/theopenlane/dbx/internal/httpserve/server"
	"github.com/theopenlane/dbx/internal/httpserve/serveropts"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start the server",
	RunE: func(cmd *cobra.Command, args []string) error {
		return serve(cmd.Context())
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.PersistentFlags().String("config", "./config/.config.yaml", "config file location")
	viperBindFlag("config", serveCmd.PersistentFlags().Lookup("config"))
}

func serve(ctx context.Context) error {
	// setup db connection for server
	var (
		err error
	)

	serverOpts := []serveropts.ServerOption{}
	serverOpts = append(serverOpts,
		serveropts.WithConfigProvider(&config.ConfigProviderWithRefresh{}),
		serveropts.WithHTTPS(),
		serveropts.WithMiddleware(),
	)

	so := serveropts.NewServerOptions(serverOpts, viper.GetString("config"))

	err = otelx.NewTracer(so.Config.Settings.Tracer, appName)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize tracer")
	}

	// create ent dependency injection
	entOpts := []ent.Option{}

	if so.Config.Settings.Providers.TursoEnabled {
		tursoClient, err := turso.NewClient(so.Config.Settings.Turso)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to initialize turso client")
		}

		entOpts = append(entOpts, ent.Turso(tursoClient))
	}

	// Setup DB connection
	entdbClient, dbConfig, err := entdb.NewMultiDriverDBClient(ctx, so.Config.Settings.DB, entOpts)
	if err != nil {
		return err
	}

	defer entdbClient.Close()

	// Setup Redis connection
	redisClient := cache.New(so.Config.Settings.Redis)
	defer redisClient.Close()

	// Add Driver to the Handlers Config
	so.Config.Handler.DBClient = entdbClient

	// Add redis client to Handlers Config
	so.Config.Handler.RedisClient = redisClient

	// add ready checks
	so.AddServerOptions(
		serveropts.WithReadyChecks(dbConfig, redisClient),
	)

	srv := server.NewServer(so.Config)

	// Setup Graph API Handlers
	so.AddServerOptions(serveropts.WithGraphRoute(srv, entdbClient))

	if err := srv.StartEchoServer(ctx); err != nil {
		log.Error().Err(err).Msg("failed to run server")
	}

	return nil
}
