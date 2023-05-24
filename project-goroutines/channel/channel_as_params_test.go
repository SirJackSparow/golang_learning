package channel

import (
	"testing"
	"time"
)

func GivenMeResponse(channel chan int) {
	time.Sleep(2 * time.Second)
	channel <- 20
}

func TestChannelAsParams(t *testing.T) {
	channel := make(chan int)
	defer close(channel)

	go GivenMeResponse(channel)

	data := <-channel

	println(data)

}
