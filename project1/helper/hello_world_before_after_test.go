package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// harus bernama TestMain
func TestMain(m *testing.M) {

	print("Before test")
	m.Run()

	print("After test")

}

func TestHello(t *testing.T) {
	result := HelloWorld("boy")

	if result != "Hello" {
		assert.Equal(t, "Hello boy", result, "Result should be Hello boy")
	}
}
