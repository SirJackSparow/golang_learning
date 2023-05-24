package tabletst

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTableTest(t *testing.T) {

	tests := []struct {
		name     string
		number   int
		request  int
		expected int
	}{
		{
			name:     "ten",
			number:   10,
			request:  10,
			expected: 20,
		},
		{
			name:     "five",
			number:   5,
			request:  5,
			expected: 10,
		},
		{
			name:     "four",
			number:   4,
			request:  4,
			expected: 8,
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			result := TableTst(test.request + test.request)
			require.Equal(t, test.expected, result)
		})
	}
}
