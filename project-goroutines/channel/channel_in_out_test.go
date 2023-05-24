package channel

import (
	"testing"
	"time"
)

// hanya menerima
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Hello Worlds"
}

// hanya mengirim
func OnlyOut(channel <-chan string) {

	data := <-channel
	println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)

	close(channel)
}
