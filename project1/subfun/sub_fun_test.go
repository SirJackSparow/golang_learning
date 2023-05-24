package subfun

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSubTest(t *testing.T) {

	one, two := SubFunc()

	t.Run("one", func(t *testing.T) {
		require.Equal(t, "one", one, "Result should be one")
	})

	t.Run("onetwo", func(t *testing.T) {
		require.Equal(t, "two", two+"l", "Result should be two")
	})
}
