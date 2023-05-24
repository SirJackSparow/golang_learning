package channel

import (
	"testing"
	"time"
)

//channel tempat komunikasi secara syncronous di goroutines
//ngirim dari goroutines satu dan yang lain goroutines
//jika tidak ada yang menerima channel maka akan ngeblock / error
//pastikan ada yang mengirim dan menerima channel

func TestChannelTutNgeblock(t *testing.T) {

	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Fendy"
		println("Done")
	}()

	data := <-channel
	println(data)

	time.Sleep(5 * time.Second)

}
