package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/tabwriter"

	"github.com/TylerBrock/colorjson"
	"github.com/Yamashou/gqlgenc/clientv2"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/theopenlane/utils/cli/rows"
	"go.uber.org/zap"

	"github.com/theopenlane/dbx/pkg/dbxclient"
)

const (
	appName         = "dbx"
	defaultRootHost = "http://localhost:1337/"
	graphEndpoint   = "query"
)

var (
	cfgFile string
	Logger  *zap.SugaredLogger
)

var (
	// RootHost contains the root url for the OpenLane API
	RootHost string
	// GraphAPIHost contains the url for the OpenLane graph api
	GraphAPIHost string
)

type CLI struct {
	Client      dbxclient.Dbxclient
	Interceptor clientv2.RequestInterceptor
	AccessToken string
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   appName,
	Short: fmt.Sprintf("a %s cli", appName),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/."+appName+".yaml)")
	ViperBindFlag("config", RootCmd.PersistentFlags().Lookup("config"))

	RootCmd.PersistentFlags().StringVar(&RootHost, "host", defaultRootHost, "api host url")
	ViperBindFlag(appName+".host", RootCmd.PersistentFlags().Lookup("host"))

	RootCmd.PersistentFlags().StringP("format", "f", "table", "output format (json, table)")
	ViperBindFlag("output.format", RootCmd.PersistentFlags().Lookup("format"))

	// Logging flags
	RootCmd.PersistentFlags().Bool("debug", false, "enable debug logging")
	ViperBindFlag("logging.debug", RootCmd.PersistentFlags().Lookup("debug"))

	RootCmd.PersistentFlags().Bool("pretty", false, "enable pretty (human readable) logging output")
	ViperBindFlag("logging.pretty", RootCmd.PersistentFlags().Lookup("pretty"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".appName" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName("." + appName)
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.SetEnvPrefix(appName)
	viper.AutomaticEnv() // read in environment variables that match

	err := viper.ReadInConfig()

	GraphAPIHost = fmt.Sprintf("%s%s", RootHost, graphEndpoint)

	setupLogging()

	if err == nil {
		Logger.Infow("using config file", "file", viper.ConfigFileUsed())
	}
}

func setupLogging() {
	cfg := zap.NewProductionConfig()
	if viper.GetBool("logging.pretty") {
		cfg = zap.NewDevelopmentConfig()
	}

	if viper.GetBool("logging.debug") {
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	} else {
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	l, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	Logger = l.Sugar().With("app", appName)
	defer Logger.Sync() //nolint:errcheck
}

// ViperBindFlag provides a wrapper around the viper bindings that panics if an error occurs
func ViperBindFlag(name string, flag *pflag.Flag) {
	err := viper.BindPFlag(name, flag)
	if err != nil {
		panic(err)
	}
}

func createClient(baseURL string) (*CLI, error) {
	cli := CLI{}

	c := dbxclient.Config{
		BaseURL: baseURL,
		Debug:   viper.GetBool("logging.debug"),
	}

	i := dbxclient.WithEmptyInterceptor()
	interceptors := []clientv2.RequestInterceptor{i}

	if viper.GetBool("logging.debug") {
		interceptors = append(interceptors, dbxclient.WithLoggingInterceptor())
	}

	cli.Client = c.NewClientWithInterceptors(interceptors)
	cli.Interceptor = i

	// new client with params
	return &cli, nil
}

func GetGraphClient() (*CLI, error) {
	return createClient(GraphAPIHost)
}

func JSONPrint(s []byte) error {
	var obj map[string]interface{}

	err := json.Unmarshal(s, &obj)
	if err != nil {
		return err
	}

	f := colorjson.NewFormatter()
	f.Indent = 2

	o, err := f.Marshal(obj)
	if err != nil {
		return err
	}

	fmt.Println(string(o))

	return nil
}

// TablePrint prints a table to the console
func TablePrint(header []string, data [][]string) error {
	w := rows.NewTabRowWriter(tabwriter.NewWriter(os.Stdout, 1, 0, 4, ' ', 0)) //nolint:mnd
	defer w.(*rows.TabRowWriter).Flush()

	if err := w.Write(header); err != nil {
		return err
	}

	for _, r := range data {
		if err := w.Write(r); err != nil {
			return err
		}
	}

	return nil
}

// GetHeaders returns the name of each field in a struct
func GetHeaders(s interface{}, prefix string) []string {
	headers := []string{}
	val := reflect.Indirect(reflect.ValueOf(s))

	// ensure we have a struct otherwise this will panic
	if val.Kind() == reflect.Struct {
		for i := range val.NumField() { //nolint:typecheck // go 1.22+ allows this, linter is wrong
			if val.Type().Field(i).Type.Kind() == reflect.Struct {
				continue
			}

			headers = append(headers, fmt.Sprintf("%s%s", prefix, val.Type().Field(i).Name))
		}
	} else {
		// if the struct is a map, get the keys
		for k := range val.Interface().(map[string]interface{}) {
			headers = append(headers, fmt.Sprintf("%s%s", prefix, k))
		}
	}

	return headers
}

// GetFields returns the value of each field in a struct
func GetFields(i interface{}) (res []string) {
	v := reflect.ValueOf(i)

	// ensure we have a struct otherwise this will panic
	if v.Kind() == reflect.Struct {
		for j := range v.NumField() { //nolint:typecheck // go 1.22+ allows this, linter is wrong
			t := v.Field(j).Type()
			if t.Kind() == reflect.Struct {
				continue
			}

			var val string

			switch t.Kind() {
			case reflect.Ptr:
				val = v.Field(j).Elem().String()
			case reflect.Slice:
				val = fmt.Sprintf("%v", v.Field(j).Interface())
			default:
				val = v.Field(j).String()
			}

			res = append(res, val)
		}
	}

	return
}

// GraphResponse is the response from the graph api containing a list of edges
type GraphResponse struct {
	Edges []Edge `json:"edges"`
}

// Edge is a single edge in the graph response
type Edge struct {
	Node interface{} `json:"node"`
}

// RowsTablePrint prints a table to the console with multiple rows using a map[string]interface{} as the row data
func RowsTablePrint(resp GraphResponse) error {
	// check if there are any groups, otherwise we have nothing to print
	if len(resp.Edges) > 0 {
		rows := resp.Edges

		data := [][]string{}

		headers := GetHeaders(rows[0].Node, "")

		// get the field values using the header names as the key to ensure the order is correct
		for _, r := range rows {
			rowMap := r.Node.(map[string]interface{})
			row := []string{}

			for _, h := range headers {
				row = append(row, rowMap[h].(string))
			}

			data = append(data, row)
		}

		// print ze data
		return TablePrint(headers, data)
	}

	return nil
}

// SingleRowTablePrint prints a single row table to the console
func SingleRowTablePrint(r interface{}) error {
	// get the headers for the table for each struct
	header := GetHeaders(r, "")

	data := [][]string{}

	// get the field values for each struct
	fields := GetFields(r)

	// append the fields to the data slice
	data = append(data, fields)

	// print ze data
	return TablePrint(header, data)
}
