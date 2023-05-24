package channel

import (
	"testing"
	"time"
)

//buffer is place for save item in channel
//commonly channel must have sender and receiver with buffer we can save it to buffer first so we can use it without receiver cause will temporary save to buffer.

func TestBufferChannel(t *testing.T) {

	//capacity 3
	channel := make(chan string, 3)
	defer close(channel)

	//without goroutines
	//channel <- "One"
	//channel <- "Two"
	//channel <- "Three"
	//channel <- "it will error because the capacity just 3"

	go func() {

		channel <- "One"
		channel <- "Two"
		channel <- "Three"
	}()

	go func() {
		println(<-channel)
		println(<-channel)
		println(<-channel)
	}()

	time.Sleep(2 * time.Second)
}

func TestBufferChannelPart2(t *testing.T) {

	channel := make(chan int, 2)

	channel <- 1
	channel <- 2

	temp1, temp2 := <-channel, <-channel

	println("data 1: ", temp1)
	println("data 2: ", temp2)

	channel <- 3
	channel <- 4

	println("data 3: ", <-channel)
	println("data 4: ", <-channel)

}
