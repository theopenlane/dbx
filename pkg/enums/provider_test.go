package enums_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/theopenlane/dbx/pkg/enums"
)

func TestToDatabaseProvider(t *testing.T) {
	testCases := []struct {
		input    string
		expected enums.DatabaseProvider
	}{
		{
			input:    "local",
			expected: enums.Local,
		},
		{
			input:    "Turso",
			expected: enums.Turso,
		},
		{
			input:    "UNKNOWN",
			expected: enums.InvalidProvider,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Convert %s to DatabaseProvider", tc.input), func(t *testing.T) {
			result := enums.ToDatabaseProvider(tc.input)
			assert.Equal(t, tc.expected, *result)
		})
	}
}
