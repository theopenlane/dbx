package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const appName = "dbx"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   appName,
	Short: "A server for scheduling geographically distributed databases",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().Bool("pretty", false, "enable pretty (human readable) logging output")
	viperBindFlag("pretty", rootCmd.PersistentFlags().Lookup("pretty"))

	rootCmd.PersistentFlags().Bool("debug", false, "debug logging output")
	viperBindFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
}

// initConfig reads in flags set for server startup
// all other configuration is done by the server with koanf
// refer to the README.md for more information
func initConfig() {
	err := viper.ReadInConfig()
	if err != nil {
		log.Info().Err(err).Str("file", viper.ConfigFileUsed()).Msg("error reading config file")
	}

	setupLogging()
}

// viperBindFlag provides a wrapper around the viper bindings that panics if an error occurs
func viperBindFlag(name string, flag *pflag.Flag) {
	err := viper.BindPFlag(name, flag)
	if err != nil {
		panic(err)
	}
}

// setupLogging sets up the logging defaults for the application
func setupLogging() {
	// setup logging with time and app name
	log.Logger = zerolog.New(os.Stderr).
		With().Timestamp().
		Logger().
		With().Str("app", appName).
		Logger()

	// set the log level
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// set the log level to debug if the debug flag is set and add additional information
	if viper.GetBool("debug") {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)

		buildInfo, _ := debug.ReadBuildInfo()

		log.Logger = log.Logger.With().
			Caller().
			Int("pid", os.Getpid()).
			Str("go_version", buildInfo.GoVersion).Logger()
	}

	// pretty logging for development
	if viper.GetBool("pretty") {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: time.RFC3339,
			FormatCaller: func(i interface{}) string {
				return filepath.Base(fmt.Sprintf("%s", i))
			},
		})
	}
}
