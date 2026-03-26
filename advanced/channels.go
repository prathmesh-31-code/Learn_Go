package advanced

import (
	"fmt"
	"time"
)

// Channels are used to send and receive data between go routines

func channel() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 5)
		ch1 <- "hello"
		ch2 <- "world"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)

	case msg2 := <-ch2: //<- channel operator used to receive signals and store in msg
		fmt.Println(msg2)
	}
}
