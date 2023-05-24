package helper

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldAssert(t *testing.T) {

	result := HelloWorld("world")

	assert.Equal(t, "Hello world", result, "Result must be Hello world")
	println("Test Done !")
}

func TestHelloWorldRequire(t *testing.T) {

	result := HelloWorld("world")

	require.Equal(t, "Hello world", result, "Result must be Hello world")
	println("tidak di eksekusi karena require panggil failnow")
}

// cara skip test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run blala")
	}
}
