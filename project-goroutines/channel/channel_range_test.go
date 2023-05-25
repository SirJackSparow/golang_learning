package channel

import "testing"

// if data continue sending
// we dont know when this sender sending data end
//looping chnnel
//range

func TestChannelRanges(t *testing.T) {

	channel := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		//close channel
		close(channel)
	}()

	//if not close this below looping will not end
	for data := range channel {
		println("data ke: ", data)
	}

	println("Done")
}
