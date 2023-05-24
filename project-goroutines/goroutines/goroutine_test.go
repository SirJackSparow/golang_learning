package goroutines

import (
	"testing"
	"time"
)

func GoroutinesTut() {
	print("Goroutines ")
}

func TestCreateGoroutines(t *testing.T) {
	go GoroutinesTut()
	println("Hello")

	time.Sleep(1 * time.Second)
}
