package enums_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/theopenlane/dbx/pkg/enums"
)

func TestToToDatabaseStatus(t *testing.T) {
	testCases := []struct {
		input    string
		expected enums.DatabaseStatus
	}{
		{
			input:    "active",
			expected: enums.Active,
		},
		{
			input:    "deleted",
			expected: enums.Deleted,
		},
		{
			input:    "DELETING",
			expected: enums.Deleting,
		},
		{
			input:    "creating",
			expected: enums.Creating,
		},
		{
			input:    "UNKNOWN",
			expected: enums.InvalidStatus,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Convert %s to DatabaseStatus", tc.input), func(t *testing.T) {
			result := enums.ToDatabaseStatus(tc.input)
			assert.Equal(t, tc.expected, *result)
		})
	}
}
