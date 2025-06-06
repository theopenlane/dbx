package enums

import (
	"fmt"
	"io"
	"strings"
)

type DatabaseStatus string

var (
	Active        DatabaseStatus = "ACTIVE"
	Creating      DatabaseStatus = "CREATING"
	Deleting      DatabaseStatus = "DELETING"
	Deleted       DatabaseStatus = "DELETED"
	InvalidStatus DatabaseStatus = "INVALID"
)

// Values returns a slice of strings that represents all the possible values of the DatabaseStatus enum.
// Possible default values are "ACTIVE", "CREATING", "DELETING", and "DELETED".
func (DatabaseStatus) Values() (kinds []string) {
	for _, s := range []DatabaseStatus{Active, Creating, Deleting, Deleted} {
		kinds = append(kinds, string(s))
	}

	return
}

// String returns the DatabaseStatus as a string
func (r DatabaseStatus) String() string {
	return string(r)
}

// ToDatabaseStatus returns the database status enum based on string input
func ToDatabaseStatus(r string) *DatabaseStatus {
	switch r := strings.ToUpper(r); r {
	case Active.String():
		return &Active
	case Creating.String():
		return &Creating
	case Deleting.String():
		return &Deleting
	case Deleted.String():
		return &Deleted
	default:
		return &InvalidStatus
	}
}

// MarshalGQL implement the Marshaler interface for gqlgen
func (r DatabaseStatus) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(`"` + r.String() + `"`))
}

// UnmarshalGQL implement the Unmarshaler interface for gqlgen
func (r *DatabaseStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("wrong type for DatabaseStatus, got: %T", v) //nolint:err113
	}

	*r = DatabaseStatus(str)

	return nil
}
