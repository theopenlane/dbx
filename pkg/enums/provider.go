package enums

import (
	"fmt"
	"io"
	"strings"
)

type DatabaseProvider string

var (
	Local           DatabaseProvider = "LOCAL"
	Turso           DatabaseProvider = "TURSO"
	InvalidProvider DatabaseProvider = "INVALID"
)

// Values returns a slice of strings that represents all the possible values of the DatabaseProvider enum.
// Possible default values are "LOCAL", and "TURSO"
func (DatabaseProvider) Values() (kinds []string) {
	for _, s := range []DatabaseProvider{Local, Turso} {
		kinds = append(kinds, string(s))
	}

	return
}

// String returns the DatabaseProvider as a string
func (r DatabaseProvider) String() string {
	return string(r)
}

// ToDatabaseProvider returns the database provider enum based on string input
func ToDatabaseProvider(p string) *DatabaseProvider {
	switch p := strings.ToUpper(p); p {
	case Local.String():
		return &Local
	case Turso.String():
		return &Turso
	default:
		return &InvalidProvider
	}
}

// MarshalGQL implement the Marshaler interface for gqlgen
func (r DatabaseProvider) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(`"` + r.String() + `"`))
}

// UnmarshalGQL implement the Unmarshaler interface for gqlgen
func (r *DatabaseProvider) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("wrong type for DatabaseProvider, got: %T", v) //nolint:err113
	}

	*r = DatabaseProvider(str)

	return nil
}
